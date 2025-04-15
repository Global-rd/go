package db

import (
	"database/sql"

	"github.com/doug-martin/goqu/v9"
	"github.com/google/uuid"
)

func GetAllBooks(db *sql.DB) ([]Book, error) {
	var result []Book

	rows, err := db.Query(`SELECT * FROM bookshelf;`)
	if err != nil {
		return result, err
	}

	defer rows.Close()

	for rows.Next() {
		var book Book
		err := rows.Scan(&book.Id, &book.Title, &book.Author, &book.Published, &book.Introduction, &book.Price, &book.Stock)
		if err != nil {
			return result, err
		}
		result = append(result, book)
	}
	return result, nil
}

func GetBook(db *sql.DB, id string) ([]Book, error) {

	var result []Book

	ds := goqu.From("bookshelf").Where(goqu.C("id").Eq(id))
	expression, _, _ := ds.ToSQL()

	rows, err := db.Query(expression)
	if err != nil {
		return result, err
	}

	for rows.Next() {
		var book Book
		err := rows.Scan(&book.Id, &book.Title, &book.Author, &book.Published, &book.Introduction, &book.Price, &book.Stock)
		if err != nil {
			return result, err
		}
		result = append(result, book)
	}

	return result, nil

}

func InsertBook(db *sql.DB, book *Book) (Book, error) {

	book.Id = string(uuid.New().String())

	ds := goqu.Insert("bookshelf").Rows(
		goqu.Record{
			"id":           book.Id,
			"title":        book.Title,
			"author":       book.Author,
			"published":    book.Published,
			"introduction": book.Introduction,
			"price":        book.Price,
			"stock":        book.Stock,
		},
	)
	expression, _, _ := ds.ToSQL()

	_, err := db.Query(expression)
	if err != nil {
		return *book, err
	}

	return *book, nil
}

func UpdateBook(db *sql.DB, book *Book) (Book, error) {
	ds := goqu.From("bookshelf").
		Where(goqu.C("id").Eq(book.Id)).
		Update().
		Set(
			goqu.Record{
				"id":           book.Id,
				"title":        book.Title,
				"author":       book.Author,
				"published":    book.Published,
				"introduction": book.Introduction,
				"price":        book.Price,
				"stock":        book.Stock,
			},
		).Executor()

	if _, err := ds.Exec(); err != nil {
		return *book, err
	}

	return *book, nil
}

func DeleteBook(db *sql.DB, id string) error {
	ds := goqu.Delete("bookshelf").Where(goqu.C("id").Eq(id)).Executor()

	if _, err := ds.Exec(); err != nil {
		return err
	}

	return nil
}
