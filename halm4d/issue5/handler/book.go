package handler

import (
	"issue5/db"
	"issue5/httpresponse"
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

func (h *BookHandler) HandleBooks(w http.ResponseWriter, r *http.Request) {
	h.logger.Info("Handling books")
	books, err := h.db.GetAllBooks()
	if err != nil {
		httpresponse.WriteErrorResponse(w, httpresponse.ErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Error:      "Error fetching books",
		})
		return
	}
	h.logger.Debug("Books fetched", "books", books)

	w.Header().Set("Content-Type", "application/json")
	httpresponse.WriteResponseBody(w, http.StatusOK, books)
	h.logger.Debug("Books encoded", "books", books)
}

func (h *BookHandler) HandleBook(w http.ResponseWriter, r *http.Request) {
	h.logger.Info("Handling book")
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
	h.logger.Debug("Book encoded", "book", book)
}
