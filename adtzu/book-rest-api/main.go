package main

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"strconv"
	"strings"
)

// Book represents a book with an ID and title.
type Book struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

// BookService handles book-related operations.
type BookService struct {
	books  []Book
	logger *slog.Logger
}

func NewBookService() *BookService {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	return &BookService{
		books: []Book{
			{ID: 1, Title: "Book One"},
			{ID: 2, Title: "Book Two"},
		},
		logger: logger,
	}
}

func (s *BookService) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		if r.URL.Path == "/books" {
			s.handleGetAllBooks(w, r)
		} else if id := s.extractID(r); id != 0 {
			s.handleGetBook(w, r, id)
		} else {
			http.Error(w, "Invalid request", http.StatusBadRequest)
		}
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (s *BookService) handleGetAllBooks(w http.ResponseWriter, r *http.Request) {
	s.logger.Info("Getting all books")
	json.NewEncoder(w).Encode(s.books)
}

func (s *BookService) handleGetBook(w http.ResponseWriter, r *http.Request, id int) {
	s.logger.Info("Getting book by ID", slog.Int("id", id))
	for _, book := range s.books {
		if book.ID == id {
			json.NewEncoder(w).Encode(book)
			return
		}
	}
	http.Error(w, "Book not found", http.StatusNotFound)
}

func (s *BookService) extractID(r *http.Request) int {
	parts := strings.Split(r.URL.Path, "/")
	if len(parts) < 3 || parts[1] != "books" {
		return 0
	}
	id, err := strconv.Atoi(parts[2])
	if err != nil {
		return 0
	}
	return id
}

func main() {
	service := NewBookService()
	http.Handle("/", service)
	fmt.Println("Server is running on port 8000")
	http.ListenAndServe(":8000", nil)
}
