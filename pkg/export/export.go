package export

import "fmt"

func ParseExportCommand(args []string) {
	if len(args) < 2 {
		fmt.Println("Args UI")
	} else {
		process_args := args[1:]
		if process_args[0] == "all" || process_args[0] == "a" {
			fmt.Println("exporting all")
		} else if process_args[0] == "shelf" || process_args[0] == "s" {
			fmt.Println("exporting shelf")
		} else if process_args[0] == "book" || process_args[0] == "b" {
			fmt.Println("exporting book")
		} else {
			fmt.Println("Unrecognized Export")
		}
	}
}
