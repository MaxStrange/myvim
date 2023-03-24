package uart

import (
	"fmt"
	"log"
	"os"
	"termite/internal/util"
	"time"

	"go.bug.st/serial"
)

// The timeout for reading from UART
const timeoutMs = 10.0

// A wrapper around the serial.Port interface.
type UartBridge struct {
	LineEnding     string
	Baud           int
	RemoveEcho     bool
	port           serial.Port
	userBuffer     chan string
	cmdChannel     chan util.Command
	bufLen         int
	readBuffer     []byte
	lastWrite      []byte
	echoIndexWrite int
	echoIndexRead  int
}

// Instantiate a UartBridge struct with the given options.
func New(lineEnding string, baud int, removeEcho bool, port serial.Port, userBuffer chan string, cmdChannel chan util.Command) *UartBridge {
	u := new(UartBridge)
	u.LineEnding = lineEnding
	u.Baud = baud
	u.RemoveEcho = removeEcho
	u.port = port
	u.userBuffer = userBuffer
	u.cmdChannel = cmdChannel
	u.bufLen = calculateBufLen(u.Baud)
	u.readBuffer = make([]byte, u.bufLen)
	u.echoIndexWrite = -1
	u.echoIndexRead = 0

	// Initialize some additional state
	setTimeout(u.port)

	return u
}

// Run until user asks to exit via the command channel.
func (ub *UartBridge) Run() {
	// Alternate reading from UART, user buffer, and control buffer
	exit := false
	for !exit {
		ub.readFromSerialPort()
		select {
		case userInput := <-ub.userBuffer:
			ub.lastWrite = []byte(userInput + ub.LineEnding)
			ub.writeToSerialPort(string(ub.lastWrite))
		case cmd := <-ub.cmdChannel:
			exit = handleCommand(cmd)
		default:
			// No user data in the buffer
		}
	}

	// Close out the serial port and exit the program cleanly
	ub.port.Close()
	os.Exit(0)
}

func calculateBufLen(baud int) int {
	// Calculate the appropriate size of the buffer so that we do not block for too long
	// 1 byte per 8 bits * 1 sec per 1000 ms * N ms * bits/second
	bufLen := int((1.0 / 8.0) * (1.0 / 1000.0) * timeoutMs * float64(baud))
	if bufLen < 1 {
		log.Fatal("Internal logic error. TimeoutMs should be larger, since the resulting bufLen is less than 1.")
	}

	return bufLen
}

func setTimeout(port serial.Port) {
	// Parse the timeoutMs into a Duration
	timeoutDuration, err := time.ParseDuration(fmt.Sprintf("%vms", timeoutMs))
	if err != nil {
		log.Fatal(err)
	}

	// Set it
	err = port.SetReadTimeout(timeoutDuration)
	if err != nil {
		log.Fatal(err)
	}
}

// Returns the number of bytes stripped off (from the left).
func (ub *UartBridge) stripEcho(nbytesRead int) int {
	// If write index is -1, we haven't sent anything yet, so there is no echo
	if ub.echoIndexWrite == -1 {
		return 0
	}

	// Nothing to do
	if nbytesRead == 0 {
		return 0
	}

	// Have already finished stripping this latest echo
	if ub.echoIndexWrite >= len(ub.lastWrite) {
		return 0
	}

	// If echoIndex is pointing to a character in lastWrite that matches the
	// character in readBuffer at the same index,
	// remove it from readBuffer and increment echoIndex
	start := 0
	end := -1
	if ub.lastWrite[ub.echoIndexWrite] != ub.readBuffer[ub.echoIndexRead] {
		// Nothing to strip
		ub.echoIndexRead = 0
		return 0
	}

	for ub.lastWrite[ub.echoIndexWrite] == ub.readBuffer[ub.echoIndexRead] {
		end = ub.echoIndexRead
		ub.echoIndexWrite += 1
		ub.echoIndexRead += 1

		if ((end - start) + 1) >= nbytesRead {
			// We have examined all the bytes that we read this time around
			ub.echoIndexRead = 0
			break
		}

		if ub.echoIndexRead >= len(ub.readBuffer) {
			// We are all done. Last user input was longer than our read buffer.
			ub.echoIndexRead = 0
			break
		}

		if ub.echoIndexWrite >= len(ub.lastWrite) {
			// We are all done. We found all the characters from the echo
			ub.echoIndexWrite = 0
			break
		}
	}

	// Sanity check
	if (end < 0) || (start < 0) || (end < start) {
		log.Fatalf("Somehow managed to compute an invalid start or end index: Start: %v, End: %v", start, end)
	}

	// Remove the subslice [start:end] from readBuffer
	ub.readBuffer = ub.readBuffer[end+1:]
	return (end - start) + 1
}

func (ub *UartBridge) readFromSerialPort() {
	// Read from serial port
	n, err := ub.port.Read(ub.readBuffer)
	if err != nil {
		log.Fatal(err)
	}

	// Strip off echo
	if ub.RemoveEcho {
		n -= ub.stripEcho(n)
	}

	// Write to console
	if n > 0 {
		output := string(ub.readBuffer[:n])
		fmt.Printf("%v", output)
	}
}

func (ub *UartBridge) writeToSerialPort(userInput string) {
	ub.echoIndexWrite = 0
	userInputAsBytes := []byte(userInput)
	n, err := ub.port.Write(userInputAsBytes)
	if err != nil {
		log.Fatal(err)
	}
	if n != len(userInputAsBytes) {
		fmt.Printf(util.Red+"<TERMITE>: Sent %v bytes, but should have sent %v\n"+util.Reset, n, len(userInputAsBytes))
	}
}

// Handles the incoming command, returning whether we should exit or not.
func handleCommand(cmd util.Command) bool {
	switch {
	case cmd.Type == util.Exit:
		return true
	default:
		return false
	}
}
