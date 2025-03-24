package routes

import (
	"advrest/db"
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
}

func (b BookController) ListBooks(w http.ResponseWriter, r *http.Request) {
	result, err := db.GetAllBooks(DbConnection)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func (b BookController) GetBooks(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	result, err := db.GetBook(DbConnection, id)
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

	result, err := db.InsertBook(DbConnection, &new_book)
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

	response, err := db.UpdateBook(DbConnection, &new_book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	/*
		ds := From("bookshelf").
			Where(C("id").Eq(id)).
			Update().
			Set(
				Record{
					"id":           new_book.Id,
					"title":        new_book.Title,
					"author":       new_book.Author,
					"published":    new_book.Published,
					"introduction": new_book.Introduction,
					"price":        new_book.Price,
					"stock":        new_book.Stock,
				},
			)
		expression, _, _ := ds.ToSQL()

		_, err = DbConnection.Query(expression)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		response := BaseResponse{
			Status:  1,
			Message: fmt.Sprintf("Book %s updated", id),
		}
	*/
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (b BookController) DeleteBook(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	err := db.DeleteBook(DbConnection, id)
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

func BookRoutes() chi.Router {
	r := chi.NewRouter()
	bookHandler := BookController{}
	r.Get("/", bookHandler.ListBooks)
	r.Post("/", bookHandler.CreateBook)
	r.Get("/{id}", bookHandler.GetBooks)
	r.Put("/{id}", bookHandler.UpdateBook)
	r.Delete("/{id}", bookHandler.DeleteBook)
	return r
}
