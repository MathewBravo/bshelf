package shelf

import (
	"github.com/MathewBravo/bshelf/pkg/book"
)

type Shelves struct {
	Shelves []Shelf
}

type Shelf struct {
	Name  string
	Books []book.Book
}

func (sh *Shelf) displayBooks() {
	// !TODO: Implement displayBooks
}

func (sh *Shelf) addBook(book book.Book) {
	sh.Books = append(sh.Books, book)
}

func (sh *Shelf) deleteBook(book book.Book) {
	for i, b := range sh.Books {
		if b.Title == book.Title && b.Author == book.Author {
			sh.Books = append(sh.Books[:i], sh.Books[i+1:]...)
			return
		}
	}
}
