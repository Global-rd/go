package routes

import (
	"advrest/db"
	"encoding/json"
	"log"
	"net/http"

	//. "github.com/doug-martin/goqu/v9"
	"github.com/go-chi/chi/v5"
)

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

		rows.Scan(
			&book.Id,
			&book.Author,
			&book.Introduction,
			&book.Price,
			&book.Published,
			&book.Stock,
			&book.Title,
		)

		result = append(result, book)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func (b BookController) GetBooks(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	log.Println("ID: ", id)
	w.Write([]byte("Get by id"))
}

func (b BookController) CreateBook(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create book"))
}

func (b BookController) UpdateBook(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	log.Println("ID: ", id)
	w.Write([]byte("Update book"))
}

func (b BookController) DeleteBook(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	log.Println("ID: ", id)
	w.Write([]byte("Delete book"))
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
