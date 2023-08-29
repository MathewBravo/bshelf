package help

import "fmt"

func PrintHelp() {
	fmt.Print("\n")
	fmt.Print("B-Shelf Commands and Options\n\n")
	fmt.Print("* indicates available [COMMAND] -help for command specific help\n")
	fmt.Print("----------------------------\n")
	fmt.Print("init\t\t\t\tInitialize a B-Shelf locally.\n\n")
	fmt.Print("shelves/svs\t\t\tProvides a list of all shelves\n\n")
	fmt.Print("set\t\t\t\tUse with [SET COMMAND] to set the state of a book\n")
	fmt.Print("\t\t\t\tValid states include: Read(r), In-Progress(ip), Abandoned(a), Not-Started(ns)\n\n")
	fmt.Print("find-book/fb\t\t\t*Used with [BOOK NAME] provides an interactive list of all books including that name\n\n")
	fmt.Print("add-book/add/ab\t\t\t*Begins the process of adding a book\n\n")
	fmt.Print("export/ex\t\t\t*Used with [EXPORT COMMAND] to initiate the export process\n")
	fmt.Print("----------------------------\n")
}
