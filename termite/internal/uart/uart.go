package uart

import (
	"fmt"
	"log"
	"time"

	"go.bug.st/serial"
)

func UartBridge(port serial.Port, lineEndingMode string, baud int, userBuffer chan string) {
	// Convert lineEndingMode to an actual line ending
	lineEnding := ""
	switch {
	case lineEndingMode == "LF":
		lineEnding = "\n"
	case lineEndingMode == "LFCR":
		lineEnding = "\n\r"
	case lineEndingMode == "CR":
		lineEnding = "\r"
	default:
		log.Fatal("Assertion failed. LineEnding is something unexpected. This is an internal logic error.")
	}

	// Set the read timeout so that our reads are non-blocking (after the specified timeout)
	//// Parse the timeoutMs into a Duration
	timeoutMs := 10.0
	timeoutDuration, err := time.ParseDuration(fmt.Sprintf("%vms", timeoutMs))
	if err != nil {
		log.Fatal(err)
	}
	err = port.SetReadTimeout(timeoutDuration)
	if err != nil {
		log.Fatal(err)
	}
	//// Calculate the appropriate size of the buffer so that we do not block for too long
	//// 1 byte per 8 bits * 1 sec per 1000 ms * N ms * bits/second
	bufLen := int((1.0 / 8.0) * (1.0 / 1000.0) * timeoutMs * float64(baud))
	if bufLen < 1 {
		log.Fatal("Internal logic error. TimeoutMs should be larger, since the resulting bufLen is less than 1.")
	}

	// Alternate reading from UART and user buffer
	readBuffer := make([]byte, bufLen)
	for true {
		// Read from serial port
		n, err := port.Read(readBuffer)
		if err != nil {
			log.Fatal(err)
		}

		// Write to console
		if n != 0 {
			// TODO
			//		-> Make sure to handle color codes and other special ASCII characters appropriately
			fmt.Printf("%v", string(readBuffer[:n]))
		}

		// Check the user buffer (nonblocking)
		select {
		case userInput := <-userBuffer:
			// Add line ending and send
			userInput += lineEnding
			userInputAsBytes := []byte(userInput)
			n, err := port.Write(userInputAsBytes)
			if err != nil {
				log.Fatal(err)
			}
			if n != len(userInputAsBytes) {
				fmt.Printf("Sent %v bytes, but should have sent %v\n", n, len(userInputAsBytes))
			}
		default:
			// No user data in the buffer
		}
	}
}
