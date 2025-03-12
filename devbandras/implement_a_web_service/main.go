package main

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"os"
)

const defaultBookFileName = "books.json"

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
	Books  []Book
	Logger *slog.Logger
}

// loadBooksFromFile beolvassa a könyveket a megadott JSON fájlból, és visszaadja őket egy Books tömbben.
// Ha a fájl nem létezik, vagy a JSON feldolgozása sikertelen, hibát ad vissza.
//
// Parameters:
// - fileName: A beolvasandó JSON fájl neve.
//
// Returns:
// - []Book: Könyvek tömbje.
// - error: Hiba esetén a hibaüzenetet tartalmazó érték.
func loadBooksFromFile(fileName string) ([]Book, error) {
	// könyv adatbázis létezik-e
	_, err := os.Stat(fileName)
	if os.IsNotExist(err) {
		return nil, fmt.Errorf("az adatbázis nem létezik: %s", fileName)
	}

	// adatbázis megnyitása
	data, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	// json feldolgozása
	var books []Book
	err = json.Unmarshal(data, &books)
	if err != nil {
		return nil, err
	}

	return books, nil
}

// A NewBookStore létrehozza a BookStore új példányát.
//
// Parameters:
// - bookFileName: A könyveket tartalmazó JSON fájl.
// - logger: slog.Logger naplózó pélánya
//
// Returns:
// - *BookStore: A létrehozott Bookstore mutatója.
// - error: Hiba esetén a hibaüzenetet tartalmazó érték.
func NewBookStore(bookFileName string, logger *slog.Logger) (*BookStore, error) {
	books, err := loadBooksFromFile(bookFileName)
	if err != nil {
		return nil, err
	}

	return &BookStore{
		Books:  books,
		Logger: logger,
	}, nil
}

func main() {
	// logger létrehozása
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}))

	// bookStore létrehozása
	bookStore, err := NewBookStore(defaultBookFileName, logger)
	if err != nil {
		logger.Error(fmt.Sprintf("Hiba a book store inicializálása közben: %s", defaultBookFileName), "error", err)
		// ha nincs adatbázi kilépünk
		os.Exit(1)
	}

	fmt.Println(bookStore)

	// ...

}
