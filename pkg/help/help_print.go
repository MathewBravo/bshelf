package help

import "fmt"

func PrintHelp() {
	fmt.Print("\n")
	fmt.Print("B-Shelf Commands and Options\n")
	fmt.Print("----------------------------\n")
	fmt.Print("\nCommands\n")
	fmt.Print("shelves svs\t\t\tProvides a list of all shelves\n")
	fmt.Print("set\t\t\t\tUse with [SET COMMAND] to set the state of a book\n")

	fmt.Print("\n")
	fmt.Print("----------------------------\n")
	fmt.Print("\nSet Commands\n")
}
