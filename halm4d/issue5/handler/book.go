package handler

import (
	"encoding/json"
	"errors"
	"issue5/db"
	"log/slog"
	"net/http"
)

type Option func(handler *BookHandler)

type BookHandler struct {
	logger *slog.Logger
	db     db.BookRepository
}

func NewBookHandler(opts ...Option) *BookHandler {
	bh := &BookHandler{}
	for _, opt := range opts {
		opt(bh)
	}
	if bh.logger == nil {
		bh.logger = slog.Default()
		bh.logger.Info("No logger provided, using default logger")
	}
	if bh.db == nil {
		bh.db = db.NewInMemoryBookRepository()
		bh.logger.Error("No database provided, using default in-memory database")
	}
	return bh
}

func WithLogger(logger *slog.Logger) Option {
	return func(handler *BookHandler) {
		handler.logger = logger
	}
}

func WithDatabase(db db.BookRepository) Option {
	return func(handler *BookHandler) {
		handler.db = db
	}
}

func (h *BookHandler) HandleBooks(w http.ResponseWriter, r *http.Request) {
	h.logger.Info("Handling books")
	books, err := h.db.GetAllBooks()
	if err != nil {
		http.Error(w, "Error fetching books", http.StatusInternalServerError)
		return
	}
	h.logger.Debug("Books fetched", "books", books)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(books)
	if err != nil {
		http.Error(w, "Error encoding books", http.StatusInternalServerError)
		return
	}
	h.logger.Debug("Books encoded", "books", books)
}

func (h *BookHandler) HandleBook(w http.ResponseWriter, r *http.Request) {
	h.logger.Info("Handling book")
	id := r.PathValue("id")
	if id == "" {
		encodeErr := json.NewEncoder(w).Encode(ErrorResponse{Error: "missing ID path variable"})
		if encodeErr != nil {
			http.Error(w, "Error encoding error response", http.StatusInternalServerError)
			return
		}
		return
	}
	w.Header().Set("Content-Type", "application/json")

	book, err := h.db.GetBookByID(id)
	if err != nil {
		switch {
		case errors.Is(err, db.NotFoundError):
			w.WriteHeader(http.StatusNotFound)
			encodeErr := json.NewEncoder(w).Encode(ErrorResponse{Error: err.Error()})
			if encodeErr != nil {
				http.Error(w, "Error encoding error response", http.StatusInternalServerError)
				return
			}
			return
		default:
			w.WriteHeader(http.StatusInternalServerError)
			encodeErr := json.NewEncoder(w).Encode(ErrorResponse{Error: err.Error()})
			if encodeErr != nil {
				http.Error(w, "Error encoding error response", http.StatusInternalServerError)
				return
			}
			return
		}
	}
	h.logger.Debug("Book fetched", "book", book)

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(book)
	if err != nil {
		http.Error(w, "Error encoding book", http.StatusInternalServerError)
		return
	}
	h.logger.Debug("Book encoded", "book", book)
}

type ErrorResponse struct {
	Error string `json:"error"`
}
