package input

import (
	"bufio"
	"log"
	"os"
	"strings"
)

// Get user input until we are given a stop command.
func GetUserInput(userBuffer chan string) {
	reader := bufio.NewReader(os.Stdin)
	for true {
		userInput, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		// Strip the final newline and optional carriage return
		userInput = strings.TrimRight(userInput, "\n\r")

		// If userInput is Ctrl+a followed by some stuff, we handle it with a special case

		// Otherwise, place it on the userBuffer channel for UART consumption
		userBuffer <- userInput
	}
	// TODO:
	// Thread for typing. Constantly read in values from the user and send them to the user buffer.
	//		-> If the user types Ctrl+a, we enter command mode and all subsequent characters (until enter) are meant for us.
	//		-> If the user types Ctrl+c or other signal commands, we need to make sure that we send them to the remote, unless we are in command mode.
}
