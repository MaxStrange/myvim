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
	LineEnding string
	Baud       int
	RemoveEcho bool
	port       serial.Port
	userBuffer chan string
	cmdChannel chan util.Command
	bufLen     int
	readBuffer []byte
	lastWrite  []byte
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

func (ub *UartBridge) readFromSerialPort() {
	// Read from serial port
	n, err := ub.port.Read(ub.readBuffer)
	if err != nil {
		log.Fatal(err)
	}

	// Strip off echo
	if ub.RemoveEcho {
		// TODO
	}

	// Write to console
	if n != 0 {
		output := string(ub.readBuffer[:n])
		fmt.Printf("%v", output)
	}
}

func (ub *UartBridge) writeToSerialPort(userInput string) {
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
