package service

import (
	"encoding/json"
	"fmt"
	"os"
)

// Egy könyv adatai
type Book struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Year     int    `json:"year"`
	Category string `json:"category"`
	ISBN     string `json:"isbn"`
}

type BookStore struct {
	Books []Book
}

// A NewBookStore létrehozza a BookStore új példányát.
//
// Parameters:
// -
//
// Returns:
// - *BookStore: A létrehozott Bookstore mutatója.
func NewBookStore() *BookStore {
	return &BookStore{}
}

// LoadBooksFromFile beolvassa a könyveket a megadott JSON fájlból, és visszaadja őket egy Books tömbben.
// Ha a fájl nem létezik, vagy a JSON feldolgozása sikertelen, hibát ad vissza.
//
// Parameters:
// - fileName: A beolvasandó JSON fájl neve.
//
// Returns:
// - []Book: Könyvek tömbje.
// - error: Hiba esetén a hibaüzenetet tartalmazó érték.
func (bs *BookStore) LoadBooksFromFile(fileName string) error {
	// könyv adatbázis létezik-e
	_, err := os.Stat(fileName)
	if os.IsNotExist(err) {
		return fmt.Errorf("az adatbázis nem létezik: %s", fileName)
	}

	// adatbázis megnyitása
	data, err := os.ReadFile(fileName)
	if err != nil {
		return err
	}

	// json feldolgozása
	var books []Book
	err = json.Unmarshal(data, &books)
	if err != nil {
		return err
	}
	bs.Books = books

	return nil
}

// Összes könyv lekérése
//
// Returns:
// - []Book: Az összes könyvet tartalmazó slice
func (bs *BookStore) GetAllBooks() []Book {
	return bs.Books
}

// Egy könyv lekérése id alapján
// Parameters:
// - id: A könyv egyedi azonosítója
//
// Returns:
// - *Book: A könyv mutatója
// - error: Hiba esetén a hibaüzenetet tartalmazó érték
func (bs *BookStore) GetBookByID(id int) (*Book, error) {
	for _, book := range bs.Books {
		if book.ID == id {
			return &book, nil
		}
	}
	return &Book{}, fmt.Errorf("Book (id: %d) not found", id)
}
