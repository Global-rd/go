package book

import (
	"database/sql"
	"log/slog"
	"main/container"
	"strconv"

	"github.com/huandu/go-sqlbuilder"
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

	queryBuilder := sqlbuilder.NewSelectBuilder()
	queryBuilder.Select("id", "date", "genre", "isbn", "writer", "title")
	queryBuilder.From("library.book")

	query := queryBuilder.String()

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

	queryBuilder := sqlbuilder.NewSelectBuilder()
	queryBuilder.SetFlavor(sqlbuilder.PostgreSQL)
	queryBuilder.Select("id", "date", "genre", "isbn", "writer", "title")
	queryBuilder.From("library.book")
	queryBuilder.Where(queryBuilder.Equal("id", bookId))

	query, args := queryBuilder.Build()
	row := c.container.GetDB().QueryRow(query, args...)

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

	queryBuilder := sqlbuilder.NewDeleteBuilder()
	queryBuilder.SetFlavor(sqlbuilder.PostgreSQL)
	queryBuilder.DeleteFrom("library.book")
	queryBuilder.Where(queryBuilder.Equal("id", bookId))

	query, args := queryBuilder.Build()
	returnString := "Book deleted successfully"

	result, err := c.container.GetDB().Exec(query, args...)
	if err != nil {
		slog.Error("Error executing delete query", slog.String("err", err.Error()))
		returnString = "Error deleting book"
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

	queryBuilder := sqlbuilder.NewInsertBuilder()
	queryBuilder.SetFlavor(sqlbuilder.PostgreSQL)
	queryBuilder.InsertInto("library.book")
	queryBuilder.Cols("date", "genre", "isbn", "title", "writer")
	queryBuilder.Values(newBook.Date, newBook.Genre, newBook.ISBN, newBook.Title, newBook.Writer)

	queryBuilder.Returning("id")

	var insertedId string

	query, args := queryBuilder.Build()
	err := c.container.GetDB().QueryRow(query, args...).Scan(&insertedId)

	if err != nil {
		slog.Error("Error executing Insert query", slog.String("err", err.Error()))
	}

	book := c.HandleBookById(insertedId)

	return book

}

func (c Controller) HandleUpdateBook(updateableBook Book, id string) Book {
	bookId, err := strconv.Atoi(id)

	if err != nil {
		slog.Error("Error decoding request body", slog.String("err", err.Error()))
	}

	queryBuilder := sqlbuilder.NewUpdateBuilder()
	queryBuilder.SetFlavor(sqlbuilder.PostgreSQL)
	queryBuilder.Update("library.book")
	queryBuilder.Set(
		queryBuilder.Assign("date", updateableBook.Date),
		queryBuilder.Assign("genre", updateableBook.Genre),
		queryBuilder.Assign("isbn", updateableBook.ISBN),
		queryBuilder.Assign("title", updateableBook.Title),
		queryBuilder.Assign("writer", updateableBook.Writer),
	)
	queryBuilder.Where(queryBuilder.Equal("id", bookId))

	query, args := queryBuilder.Build()

	var insertedId string

	err = c.container.GetDB().QueryRow(query, args...).Scan(&insertedId)

	if err != nil {
		slog.Error("Error during update", slog.String("err", err.Error()))
	}

	book := c.HandleBookById(id)

	return book
}
