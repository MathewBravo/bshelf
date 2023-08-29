package init_shelf

import (
	"database/sql"
	"io/ioutil"
	"log"
	"os"
	"os/user"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
)

func Init() {
	// Create bshelf directory
	// Create shelves table sqlite
	// Create books table sqlite
	// Create notes table sqlite
	ok := createBShelfDirectory()
	if !ok {
		return
	}
	ok = createDatabase()
	if !ok {
		return
	}
}

func createBShelfDirectory() bool {
	curUser, err := user.Current()
	if err != nil {
		log.Fatal(err)
		return false
	}

	bshelfPath := filepath.Join(curUser.HomeDir, ".bshelf")

	err = os.Mkdir(bshelfPath, 0755)
	if err != nil {
		if os.IsExist(err) {
			log.Fatal("B-Shelf directory already exists")
			return false
		} else {
			log.Fatal(err)
			return false
		}
	} else {
		log.Printf("Created B-Shelf directory at %s", bshelfPath)
	}
	return true
}

func createDatabase() bool {
	curUser, err := user.Current()
	if err != nil {
		log.Fatal(err)
		return false
	}

	bshelfPath := filepath.Join(curUser.HomeDir, ".bshelf")
	dbPath := filepath.Join(bshelfPath, "bookshelf.db")
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		log.Fatal(err)
		return false
	}
	defer db.Close()

	_, err = db.Exec(`CREATE TABLE Shelves (
    ShelfID INTEGER PRIMARY KEY AUTOINCREMENT,
    Name TEXT NOT NULL
	);`)
	if err != nil {
		log.Fatal(err)
		return false
	}

	_, err = db.Exec(`CREATE TABLE Books (
	BookID INTEGER PRIMARY KEY AUTOINCREMENT,
    Title TEXT NOT NULL,
    Author TEXT,
    State TEXT NOT NULL DEFAULT 'Not-Started'
	);`)
	if err != nil {
		log.Fatal(err)
		return false
	}

	_, err = db.Exec(`CREATE TABLE BookShelves (
    ShelfID INTEGER,
    BookID INTEGER,
    PRIMARY KEY (ShelfID, BookID),
    FOREIGN KEY (ShelfID) REFERENCES Shelves(ShelfID),
    FOREIGN KEY (BookID) REFERENCES Books(BookID)
	);`)
	if err != nil {
		log.Fatal(err)
		return false
	}

	_, err = db.Exec(`CREATE TABLE Notes (
		    NoteID INTEGER PRIMARY KEY AUTOINCREMENT,
    BookID INTEGER,
    Content TEXT NOT NULL,
    FOREIGN KEY (BookID) REFERENCES Books(BookID)
	);`)

	if err != nil {
		log.Fatal(err)
		return false
	}

	files, err := ioutil.ReadDir(bshelfPath)
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		log.Println(file.Name())
	}

	return true
}
