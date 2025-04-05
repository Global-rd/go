package book

import (
	"database/sql"
	"fmt"
	"log/slog"
	"main/container"
	"strconv"
)

type Book struct {
	ID     int
	Title  string
	Writer string
	Genre  string
	Date   string
	ISBN   string
}

type Controller struct {
	container container.Container
}

func NewController(container container.Container) Controller {
	return Controller{
		container: container,
	}
}

func (c Controller) IsItWorks() string {
	ret := "It Works!"
	return ret
}

func (c Controller) HandleBooks() ([]Book, error) {
	var books []Book

	query := "SELECT id, date, genre, isbn, writer, title FROM library.book"
	rows, err := c.container.GetDB().Query(query)
	for rows.Next() {
		var book Book
		err = rows.Scan(&book.ID, &book.Date, &book.Genre, &book.ISBN, &book.Title, &book.Writer)
		books = append(books, book)
	}

	if err != nil {
		slog.Error(err.Error())
	}

	return books, nil
}

func (c Controller) HandleBookById(id string) Book {
	bookId, err := strconv.Atoi(id)

	if err != nil {
		slog.Error(err.Error())
	}

	var book Book

	query := "SELECT id, date, genre, isbn, writer, title FROM library.book WHERE id = $1"
	row := c.container.GetDB().QueryRow(query, bookId)

	err = row.Scan(&book.ID, &book.Date, &book.Genre, &book.ISBN, &book.Writer, &book.Title)
	if err != nil {
		if err == sql.ErrNoRows {
			slog.Error("Error selecting book", slog.String("err", err.Error()))
		}
	}

	if err != nil {
		slog.Error(err.Error())
	}

	return book

}

func (c Controller) HandleDeleteBookById(id string) string {
	bookId, err := strconv.Atoi(id)

	if err != nil {
		slog.Error(err.Error())
	}

	returnString := "Book deleted successfully"

	query := "DELETE FROM library.book WHERE id = $1"
	result, err := c.container.GetDB().Exec(query, bookId)

	if err != nil {
		if err == sql.ErrNoRows {
			returnString = "Error deleting"
		}
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		returnString = "Error deleting - Unable to determine rows affected"
	} else if rowsAffected == 0 {
		returnString = "Error deleting - No book was deleted with the given ID"
	}

	return returnString

}

func (c Controller) HandleCreateBook(newBook Book) Book {

	fmt.Println(newBook)
	query := "INSERT INTO library.boo (date, genre, isbn, title, writer, created_at) VALUES($1, $2, $3, $4, $5, now())"
	result, err := c.container.GetDB().Exec(query, newBook.Date, newBook.Genre, newBook.ISBN, newBook.Title, newBook.Writer)
	fmt.Println(result)
	if err != nil {
		if err == sql.ErrNoRows {
			slog.Error("Error inserting book", slog.String("err", err.Error()))
		}
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		slog.Error("Error retrieving rows affected", slog.String("err", err.Error()))
		return Book{}
	}
	if rowsAffected == 0 {
		slog.Error("Error inserting book - No book was inserted")
		return Book{}
	}

	return Book{}

}

func (c Controller) HandleUpdateBook(updateableBook Book, id string) Book {
	bookId, err := strconv.Atoi(id)

	if err != nil {
		slog.Error("Error decoding request body", slog.String("err", err.Error()))
	}
	fmt.Println(updateableBook)
	query := "UPDATE library.book " +
		" SET date = $1, " +
		" genre = $2, " +
		" isbn = $3, " +
		" title = $4, " +
		" writer = $5 " +
		" WHERE id = $6; "

	result, err := c.container.GetDB().Exec(query,
		updateableBook.Date, updateableBook.Genre, updateableBook.ISBN,
		updateableBook.Title, updateableBook.Writer, bookId)

	if err != nil {
		slog.Error("Error during update", slog.String("err", err.Error()))
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		slog.Error("Error retrieving rows affected", slog.String("err", err.Error()))
	} else if rowsAffected == 0 {
		slog.Error("No rows affected during update")
	}

	book := c.HandleBookById(id)

	return book
}
