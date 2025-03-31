package routes

import (
	"advrest/db"
	"advrest/logger"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type BaseResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type BookController struct {
	DbConnection *sql.DB
	Logger       *logger.Log
}

func (b BookController) ListBooks(w http.ResponseWriter, r *http.Request) {
	result, err := db.GetAllBooks(b.DbConnection)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		b.Logger.ERROR("Internal error at get all books")
		return
	}
	b.Logger.INFO("Get all books served")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func (b BookController) GetBooks(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	result, err := db.GetBook(b.DbConnection, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func (b BookController) CreateBook(w http.ResponseWriter, r *http.Request) {
	var new_book db.Book
	err := json.NewDecoder(r.Body).Decode(&new_book)
	if err != nil {
		http.Error(w, "invalid json body", http.StatusBadRequest)
		return
	}

	result, err := db.InsertBook(b.DbConnection, &new_book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func (b BookController) UpdateBook(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var new_book db.Book
	err := json.NewDecoder(r.Body).Decode(&new_book)
	if err != nil {
		http.Error(w, "invalid json body", http.StatusBadRequest)
		return
	}

	new_book.Id = id

	response, err := db.UpdateBook(b.DbConnection, &new_book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (b BookController) DeleteBook(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	err := db.DeleteBook(b.DbConnection, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := BaseResponse{
		Status:  1,
		Message: fmt.Sprintf("Book %s deleted", id),
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func BookRoutes(dbConnection *sql.DB, logger *logger.Log) chi.Router {
	r := chi.NewRouter()
	bookHandler := BookController{
		DbConnection: dbConnection,
		Logger:       logger,
	}
	r.Get("/", bookHandler.ListBooks)
	r.Post("/", bookHandler.CreateBook)
	r.Get("/{id}", bookHandler.GetBooks)
	r.Put("/{id}", bookHandler.UpdateBook)
	r.Delete("/{id}", bookHandler.DeleteBook)
	return r
}
