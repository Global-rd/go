package handler

import (
	"encoding/json"
	"github.com/google/uuid"
	"issue6/db"
	"issue6/httpresponse"
	"log/slog"
	"net/http"
)

type Option func(handler *BookHandler)

type BookHandler struct {
	logger *slog.Logger
	db     db.BookRepository
}

func NewBookHandler(opts ...Option) *BookHandler {
	bh := &BookHandler{
		logger: slog.Default(),
		db:     db.NewInMemoryBookRepository(),
	}
	for _, opt := range opts {
		opt(bh)
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

func (h *BookHandler) HandleGetAllBooks(w http.ResponseWriter, r *http.Request) {
	h.logger.Info("Handling get all books")
	books, err := h.db.GetAllBooks()
	if err != nil {
		httpresponse.WriteErrorResponse(w, httpresponse.ErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Error:      err.Error(),
		})
		return
	}
	h.logger.Debug("Books fetched", "books", books)

	w.Header().Set("Content-Type", "application/json")
	httpresponse.WriteResponseBody(w, http.StatusOK, books)
	h.logger.Debug("Book get all response sent", "books", books)
}

func (h *BookHandler) HandleGetBookById(w http.ResponseWriter, r *http.Request) {
	h.logger.Info("Handling get book by id")
	id := r.PathValue("id")
	if id == "" {
		httpresponse.WriteErrorResponse(w, httpresponse.ErrorResponse{
			StatusCode: http.StatusBadRequest,
			Error:      "missing ID path variable",
		})
		return
	}
	w.Header().Set("Content-Type", "application/json")

	book, err := h.db.GetBookByID(id)
	if errorResponse, ok := httpresponse.HandleError(err); !ok {
		httpresponse.WriteErrorResponse(w, *errorResponse)
		return
	}
	h.logger.Debug("Book fetched", "book", book)

	httpresponse.WriteResponseBody(w, http.StatusOK, book)
	h.logger.Debug("Book get by id response sent", "book", book)
}

func (h *BookHandler) HandleCreateBook(w http.ResponseWriter, r *http.Request) {
	h.logger.Info("Handling create book")
	var book db.Book
	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		httpresponse.WriteErrorResponse(w, httpresponse.ErrorResponse{
			StatusCode: http.StatusBadRequest,
			Error:      "Error decoding request body",
			Details:    err.Error(),
		})
		return
	}
	w.Header().Set("Content-Type", "application/json")

	book.ID = uuid.NewString()
	book, err := h.db.CreateBook(book)
	if errorResponse, ok := httpresponse.HandleError(err); !ok {
		httpresponse.WriteErrorResponse(w, *errorResponse)
		return
	}
	h.logger.Debug("Book created", "book", book)

	httpresponse.WriteResponseBody(w, http.StatusCreated, book)
	h.logger.Debug("Book created response sent", "book", book)
}

func (h *BookHandler) HandleUpdateBook(w http.ResponseWriter, r *http.Request) {
	h.logger.Info("Handling update book")
	id := r.PathValue("id")
	if id == "" {
		httpresponse.WriteErrorResponse(w, httpresponse.ErrorResponse{
			StatusCode: http.StatusBadRequest,
			Error:      "missing ID path variable",
		})
		return
	}
	var book db.Book
	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		httpresponse.WriteErrorResponse(w, httpresponse.ErrorResponse{
			StatusCode: http.StatusBadRequest,
			Error:      "Error decoding request body",
			Details:    err.Error(),
		})
		return
	}
	w.Header().Set("Content-Type", "application/json")

	book, err := h.db.UpdateBook(id, book)
	if errorResponse, ok := httpresponse.HandleError(err); !ok {
		httpresponse.WriteErrorResponse(w, *errorResponse)
		return
	}
	h.logger.Debug("Book updated", "book", book)

	httpresponse.WriteResponseBody(w, http.StatusOK, book)
	h.logger.Debug("Book updated response sent", "book", book)
}

func (h *BookHandler) HandleDeleteBook(w http.ResponseWriter, r *http.Request) {
	h.logger.Info("Handling delete book")
	id := r.PathValue("id")
	if id == "" {
		httpresponse.WriteErrorResponse(w, httpresponse.ErrorResponse{
			StatusCode: http.StatusBadRequest,
			Error:      "missing ID path variable",
		})
		return
	}
	w.Header().Set("Content-Type", "application/json")

	err := h.db.DeleteBook(id)
	if errorResponse, ok := httpresponse.HandleError(err); !ok {
		httpresponse.WriteErrorResponse(w, *errorResponse)
		return
	}
	h.logger.Debug("Book deleted", "id", id)

	httpresponse.WriteResponseBody(w, http.StatusNoContent, nil)
	h.logger.Debug("Book deleted response sent", "id", id)
}
