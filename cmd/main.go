package main

import (
	"os"

	"github.com/MathewBravo/bshelf/pkg/command_parsing"
)

func main() {
	args_verbose := os.Args
	args_actual := args_verbose[1:]
	cmd := command_parsing.ParseArgs(args_actual)
	command_parsing.CommandDelegator(args_actual, cmd)
}
