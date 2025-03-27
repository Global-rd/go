package db

import "slices"

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

func (r *InMemoryBookRepository) CreateBook(b Book) (Book, error) {
	r.books = append(r.books, b)
	return b, nil
}

func (r *InMemoryBookRepository) UpdateBook(id string, b Book) (Book, error) {
	for i, book := range r.books {
		if book.ID == id {
			r.books[i] = b
			return b, nil
		}
	}
	return Book{}, NotFoundError
}

func (r *InMemoryBookRepository) DeleteBook(id string) error {
	for i, book := range r.books {
		if book.ID == id {
			r.books = slices.Delete(r.books, i, i+1)
			return nil
		}
	}
	return NotFoundError
}
