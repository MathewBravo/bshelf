package add

type AddCommand int

const (
	SHELF AddCommand = iota
	BOOK
	NOTE
)

func ParseAddCommand(args []string) {

}
