package routes

import (
	"advrest/db"
	"encoding/json"
	"fmt"
	"net/http"

	. "github.com/doug-martin/goqu/v9"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

type EmptyResponse struct {
	res []db.Book `json:"books"`
}

type BaseResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

type BookController struct {
}

func (b BookController) ListBooks(w http.ResponseWriter, r *http.Request) {
	rows, err := DbConnection.Query(`SELECT * FROM bookshelf;`)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer rows.Close()

	var result []db.Book

	for rows.Next() {
		var book db.Book
		err := rows.Scan(&book.Id, &book.Title, &book.Author, &book.Published, &book.Introduction, &book.Price, &book.Stock)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		result = append(result, book)
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
	ds := From("bookshelf").Where(C("id").Eq(id))
	expression, _, _ := ds.ToSQL()

	rows, err := DbConnection.Query(expression)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var result []db.Book

	for rows.Next() {
		var book db.Book
		err := rows.Scan(&book.Id, &book.Title, &book.Author, &book.Published, &book.Introduction, &book.Price, &book.Stock)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		result = append(result, book)
	}

	if len(result) == 0 {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(EmptyResponse{})
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

	new_book.Id = string(uuid.New().String())

	ds := Insert("bookshelf").Rows(
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

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(new_book)
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
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (b BookController) DeleteBook(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	ds := Delete("bookshelf").Where(C("id").Eq(id))
	expression, _, _ := ds.ToSQL()

	_, err := DbConnection.Query(expression)
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
