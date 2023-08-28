package add

import "fmt"

type AddCommand int

const (
	SHELF AddCommand = iota
	BOOK
	NOTE
)

func ParseAddCommand(args []string) {
	switch args[0] {
	case "shelf":
		fmt.Println("ADDING SHELF")
	case "book":
		fmt.Println("ADDING BOOK")
	case "note":
		fmt.Println("ADDING NOTE")
	}
}
