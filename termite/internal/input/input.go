package input

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"os/signal"
	"strings"
	"termite/internal/util"
)

const (
	ctrlA string = "\u0001"
	ctrlC string = "\u0003"
)

func captureSignals(userBuffer chan string) {
	signalChannel := make(chan os.Signal)
	signal.Notify(signalChannel, os.Interrupt)

	for true {
		<-signalChannel
		userBuffer <- ctrlC
	}
}

func handleCommand(userInput string, cmdBuffer chan util.Command) {
	ui := strings.TrimSpace(userInput)
	ui = strings.ToLower(ui)
	ui = strings.TrimPrefix(ui, ctrlA)
	switch {
	case ui == "exit":
		cmdBuffer <- util.Command{Type: util.Exit}
	case ui == "help":
		fmt.Printf(util.Red+"<TERMITE>: %v\n"+util.Reset, util.Command{Type: util.Help})
	default:
		fmt.Printf(util.Red+"<TERMITE>: Unrecognized command: %v\n"+util.Reset, ui)
	}
}

// Get user input until we are given a stop command.
func GetUserInput(userBuffer chan string, cmdBuffer chan util.Command) {
	// TODO: Can't just read until we see a newline. We need to forward characters every so often instead of waiting for newline

	// Get OS signals like CTRL+C so we can forward them
	go captureSignals(userBuffer)

	reader := bufio.NewReader(os.Stdin)
	for true {
		userInput, err := reader.ReadString('\n')
		if err != nil && !errors.Is(err, io.EOF) {
			log.Fatal(err)
		} else if strings.HasPrefix(userInput, ctrlA) {
			// User entered CTRL+A; the rest of the text is a command
			handleCommand(userInput, cmdBuffer)
		} else {
			// Strip the final newline and optional carriage return
			userInput = strings.TrimRight(userInput, "\n\r")
			// Place it on the userBuffer channel for UART consumption
			userBuffer <- userInput
		}
	}
}
