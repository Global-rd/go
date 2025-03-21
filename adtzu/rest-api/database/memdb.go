package database

import (
	"sync"

	"rest-api/models"
)

type MemDB struct {
	books map[string]models.Book
	mu    sync.RWMutex
}

func NewMemDB() *MemDB {
	return &MemDB{
		books: make(map[string]models.Book),
	}
}

func (db *MemDB) CreateBook(book models.Book) {
	db.mu.Lock()
	db.books[book.ID] = book
	db.mu.Unlock()
}

func (db *MemDB) GetBook(id string) (models.Book, bool) {
	db.mu.RLock()
	book, ok := db.books[id]
	db.mu.RUnlock()
	return book, ok
}

func (db *MemDB) UpdateBook(id string, book models.Book) bool {
	db.mu.Lock()
	_, ok := db.books[id]
	if ok {
		db.books[id] = book
	}
	db.mu.Unlock()
	return ok
}

func (db *MemDB) DeleteBook(id string) bool {
	db.mu.Lock()
	_, ok := db.books[id]
	if ok {
		delete(db.books, id)
	}
	db.mu.Unlock()
	return ok
}
