package util

type CommandType int

const (
	// Exit the program cleanly
	Exit CommandType = iota
	Help CommandType = iota
)

const (
	HelpText string = `Available Commands:
	exit: Exit termite
	help: Print this message`
)

type Command struct {
	// Must be one of the available CommandTypes
	Type CommandType
	// Room for if we want to add args in the future
}

func (cmd Command) String() string {
	switch cmd.Type {
	case Exit:
		return "Exit"
	case Help:
		return HelpText
	default:
		return "Unrecognized Command"
	}
}
