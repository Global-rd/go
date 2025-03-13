package db

import (
	"errors"
)

type Book struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

type BookRepository interface {
	GetAllBooks() ([]Book, error)
	GetBookByID(id string) (Book, error)
}

type InMemoryBookRepository struct {
	books []Book
}

func NewInMemoryBookRepository() *InMemoryBookRepository {
	return &InMemoryBookRepository{books: []Book{
		{ID: "550e8400-e29b-41d4-a716-446655440005", Title: "The Hitchhiker's Guide to the Galaxy", Author: "Douglas Adams"},
		{ID: "550e8400-e29b-41d4-a716-446655440006", Title: "Ready Player One", Author: "Ernest Cline"},
		{ID: "550e8400-e29b-41d4-a716-446655440007", Title: "Dune", Author: "Frank Herbert"},
	}}
}

func (r *InMemoryBookRepository) GetAllBooks() ([]Book, error) {
	return r.books, nil
}

func (r *InMemoryBookRepository) GetBookByID(id string) (Book, error) {
	for _, book := range r.books {
		if book.ID == id {
			return book, nil
		}
	}
	return Book{}, NotFoundError
}

var NotFoundError = errors.New("resource not found")
