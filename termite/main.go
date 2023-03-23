package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"termite/internal/input"
	"termite/internal/uart"

	"go.bug.st/serial"
	"golang.org/x/exp/slices"
)

// Validates the given arguments and returns a sanity-checked serial.Mode object with filled-in values.
// Also validates that the given port is available.
func validateArgs(port string, baud int, dataBits int, parity string, stopBits string, lineEndings string) *serial.Mode {
	ports, err := serial.GetPortsList()
	if err != nil {
		log.Fatal(err)
	}

	// Check if the port is default ''
	if port == "" {
		log.Fatal("Either use --port to specify the port or run with --scan-ports")
	}

	// Find the port
	foundPort := false
	for _, p := range ports {
		if p == port {
			foundPort = true
			break
		}
	}
	if !foundPort {
		log.Fatal("Could not find the requested port. Try --scan-ports to see available ports.")
	}

	// Validate dataBits
	if !slices.Contains([]int{5, 6, 7, 8}, dataBits) {
		log.Fatal("Please specify an allowed --data-bits. Must be one of [5, 6, 7, 8]")
	}

	// Validate parity
	parityValue := serial.NoParity
	switch {
	case parity == "none":
		parityValue = serial.NoParity
	case parity == "odd":
		parityValue = serial.OddParity
	case parity == "even":
		parityValue = serial.EvenParity
	case parity == "mark":
		parityValue = serial.MarkParity
	case parity == "space":
		parityValue = serial.SpaceParity
	default:
		log.Fatal("Please specify an allowed --parity. Must be one of ['none', 'odd', 'even', 'mark', 'space']")
	}

	// Validate stopBits
	stopBitsValue := serial.OneStopBit
	switch {
	case stopBits == "1":
		stopBitsValue = serial.OneStopBit
	case stopBits == "1.5":
		stopBitsValue = serial.OnePointFiveStopBits
	case stopBits == "2":
		stopBitsValue = serial.TwoStopBits
	default:
		log.Fatal("Please specify an allowed --stop-bits. Must be one of [1, 1.5, 2]")
	}

	// Validate line endings
	if !slices.Contains([]string{"LF", "LFCR", "CR"}, lineEndings) {
		log.Fatal("Please specify an allowed --line-endings. Must be one of ['LF', 'LFCR', 'CR']")
	}

	mode := &serial.Mode{
		BaudRate: baud,
		DataBits: dataBits,
		Parity:   parityValue,
		StopBits: stopBitsValue,
	}
	return mode
}

// Scan ports, print them, and exit cleanly.
func scanAndExit() {
	ports, err := serial.GetPortsList()
	if err != nil {
		log.Fatal(err)
	}

	if len(ports) == 0 {
		fmt.Print("Could not find any ports.")
		os.Exit(0)
	}

	fmt.Println("Found ports:")
	for _, port := range ports {
		fmt.Printf("%v\n", port)
	}
	os.Exit(0)
}

func main() {
	var portFlag = flag.String("port", "", "Port to open")
	var baudFlag = flag.Int("baud", 115200, "Baud rate")
	var nDataBitsFlag = flag.Int("data-bits", 8, "Number of data bits. Choices: [5, 6, 7, 8]")
	var parityFlag = flag.String("parity", "none", "Parity mode. Choices: ['none', 'odd', 'even', 'mark', 'space']")
	var stopBitsFlag = flag.String("stop-bits", "1", "Number of stop bits. Choices: [1, 1.5, 2]")
	var scanPortsFlag = flag.Bool("scan-ports", false, "If given, we scan the available ports, print them, and exit")
	var lineEndingsFlag = flag.String("line-endings", "LF", "Line ending mode. Choices: [LF, LFCR, CR]")
	flag.Parse()

	// First check if the user wants to scan ports
	if *scanPortsFlag {
		scanAndExit()
	}

	options := validateArgs(*portFlag, *baudFlag, *nDataBitsFlag, *parityFlag, *stopBitsFlag, *lineEndingsFlag)
	port, err := serial.Open(*portFlag, options)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Use CTRL+a followed by x then <enter> to exit.")
	c := make(chan string)
	go uart.UartBridge(port, *lineEndingsFlag, *baudFlag, c)
	go input.GetUserInput(c)
	for true {
		// Main thread has nothing to do
	}
}
