package command_parsing

import (
	"fmt"

	"github.com/MathewBravo/bshelf/pkg/add"
	"github.com/MathewBravo/bshelf/pkg/help"
)

type Command int

const (
	SHELVES Command = iota
	FIND_BOOK
	SET
	ADD

	HELP
	UNRECOGNIZED_COMMAND
)

func ParseArgs(args []string) Command {
	switch args[0] {
	case "shelves":
		return SHELVES
	case "set":
		return SET
	case "fb", "find-book", "f-book":
		return FIND_BOOK
	case "add":
		return ADD
	case "h", "-h", "help", "-help":
		return HELP
	default:
		return UNRECOGNIZED_COMMAND
	}
}

func CommandDeligator(args []string, cmd Command) {
	switch cmd {
	case SHELVES:
		fmt.Println("SHELVES")
	case FIND_BOOK:
		fmt.Println("FIND_BOOK")
	case SET:
		fmt.Println("SET")
	case UNRECOGNIZED_COMMAND:
		fmt.Printf("The Command %s is not found", args[0])
	case ADD:
		add.ParseAddCommand(args[1:])
	case HELP:
		help.PrintHelp()
	default:
		fmt.Println("Command Not Found", args[0])
	}
}
