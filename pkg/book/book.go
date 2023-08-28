package book

type Book struct {
	Title  string
	Author string
	Notes  []Note
}

type Note struct {
	Contents string
}

func New() Book {
	return Book{}
}

func (b *Book) addNode(contents string) {
	n := Note{Contents: contents}
	b.Notes = append(b.Notes, n)
}

func (b *Book) displayNotes() {
	// TODO: Implement displayNotes
}
