package book

import (
	"fmt"

	"github.com/doug-martin/goqu/v9"
)

type BookService struct {
	DB *goqu.Database
}

func NewBookService(db *goqu.Database) *BookService {
	return &BookService{
		DB: db,
	}
}

// GetAllBooks lekéri az összes könyvet az adatbázisból
//
// Returns:
//   - []Book: tartalmazza az összes könyvet a book struktúra alapján
//   - error: nil ha nincs hiba, ellenkező esetben a hibát leíró érték
func (b *BookService) GetAllBooks() ([]Book, error) {
	var books []Book

	// dataset lekérdezése a books-ba
	dataSet := b.DB.From("app.book").Select("*")
	err := dataSet.ScanStructs(&books)
	if err != nil {
		return nil, fmt.Errorf("hiba az app.book lekérdezése során: %v", err)
	}

	return books, nil
}

// GetBookByID lekéri a megadott azonosítójú könyvet az adatbázisból.
//
// Paraméterek:
//   - id: a lekérdezni kívánt könyv azonosítója
//
// Visszatérési értékek:
//   - Book: a keresett könyv adatait tartalmazó Book struktúra
//   - error: nil, ha nincs hiba, ellenkező esetben a hibát leíró érték
func (b *BookService) GetBookByID(id int) (Book, error) {
	var book Book

	// dataset lekérdezése a book-ba
	dataSet := b.DB.From("app.book").Where(goqu.Ex{"id": id}).Select("*")
	_, err := dataSet.ScanStruct(&book)
	if err != nil {
		return book, fmt.Errorf("hiba az app.book lekérdezése során: %v", err)
	}

	return book, nil
}

// CreateBook beszúr egy új könyvet az adatbázisba.
//
// Paraméterek:
//   - book: a beszúrni kívánt könyv adatait tartalmazó Book struktúra
//
// Visszatérési értékek:
//   - int64: az újonnan beszúrt könyv azonosítója
//   - error: nil, ha nincs hiba, ellenkező esetben a hibát leíró érték
func (b *BookService) CreateBook(book Book) (int64, error) {

	// dataset beszúrásra (id-t nem adunk meg; az adatbázisban bigserial)
	dataSet := b.DB.Insert("app.book").Rows(goqu.Record{
		"title":          book.Title,
		"author":         book.Author,
		"published_year": book.PublishedYear,
		"genre":          book.Genre,
		"price":          book.Price,
	}).Returning("id")

	// lekérjük az új id-t
	var newID int64
	_, err := dataSet.Executor().ScanVal(&newID)
	if err != nil {
		return -1, fmt.Errorf("hiba az app.book beszúrásakor: %v", err)
	}

	return newID, nil
}

// UpdateBook frissíti a megadott könyv adatait az adatbázisban.
//
// Paraméterek:
//   - book: a frissíteni kívánt könyv adatait tartalmazó Book struktúra.
//
// Visszatérési értékek:
//   - error: nil, ha nincs hiba, ellenkező esetben a hibát leíró érték.
func (b *BookService) UpdateBook(book Book) error {

	// dataset frissítése
	dataSet := b.DB.Update("app.book").Set(goqu.Record{
		"title":          book.Title,
		"author":         book.Author,
		"published_year": book.PublishedYear,
		"genre":          book.Genre,
		"price":          book.Price,
	}).Where(goqu.Ex{"id": book.ID})

	_, err := dataSet.Executor().Exec()
	if err != nil {
		return fmt.Errorf("hiba az app.book frissítésekor: %v", err)
	}

	return nil
}

// DeleteBook törli a megadott azonosítójú könyvet az adatbázisból.
//
// Paraméterek:
//   - id: a törlendő könyv azonosítója
//
// Visszatérési értékek:
//   - error: nil, ha nincs hiba, ellenkező esetben a hibát leíró érték
func (b *BookService) DeleteBook(id int) error {
	// dataset törlése is alapján
	dataSet := b.DB.Delete("app.book").Where(goqu.Ex{"id": id})

	_, err := dataSet.Executor().Exec()
	if err != nil {
		return fmt.Errorf("hiba az app.book törlésekor: %v", err)
	}

	return nil
}
