package db

import (
	"errors"
)

type Book struct {
	ID     string `json:"id" db:"id"`
	Title  string `json:"title" db:"title"`
	Author string `json:"author" db:"author"`
}

type BookRepository interface {
	GetAllBooks() ([]Book, error)
	GetBookByID(id string) (Book, error)
	CreateBook(Book) (Book, error)
	UpdateBook(string, Book) (Book, error)
	DeleteBook(string) error
}

var NotFoundError = errors.New("resource not found")
var AlreadyExistsError = errors.New("resource already exists")
var FailedToCreateError = errors.New("failed to create resource")
var FailedToUpdateError = errors.New("failed to update resource")
