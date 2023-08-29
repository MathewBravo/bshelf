package book

type Book struct {
	Title  string
	Author string
	State  string
	Notes  []Note
}

type Note struct {
	Contents string
}

func New() Book {
	return Book{}
}

func (b *Book) addNote(contents string) {
	n := Note{Contents: contents}
	b.Notes = append(b.Notes, n)
}

func (b *Book) deleteNote(id int) {
	for i := range b.Notes {
		if i == id-1 {
			b.Notes = append(b.Notes[:i], b.Notes[i+1:]...)
			return
		}
	}
}

func (b *Book) displayNotes() {
	// TODO: Implement displayNotes
}
