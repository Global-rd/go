package routes

import (
	"encoding/json"
	"main/database"
	"net/http"
)

func HandleBooks(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":
		{
			getBook(w, r)
		}
	case "POST":
		{
			updateBook(w, r)
		}
	case "DELETE":
		{
			deleteBook(w, r)
		}
	case "PUT":
		{
			addBook(w, r)
		}
	default:
		http.Error(w, "invalid method", http.StatusBadRequest)
		return
	}

}

func getBook(w http.ResponseWriter, r *http.Request) {
	book_id := r.PathValue("id")

	books, err := database.DialStore()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if book_id == "" {
		w.Header().Set("Content-Type", "application/json")
		b, _ := books.FindAll()
		json.NewEncoder(w).Encode(b)
		return
	}

	book, err := books.FindOne("Id", book_id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(book)
}

func addBook(w http.ResponseWriter, r *http.Request) {}

func deleteBook(w http.ResponseWriter, r *http.Request) {
	book_id := r.PathValue("id")
	books, err := database.DialStore()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if book_id == "" {
		http.Error(w, "removing all entries is forbidden", http.StatusForbidden)
		return
	}

	err = books.DeleteOne("Id", book_id)
	if err != nil {
		http.Error(w, "book not found", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	b, _ := books.FindAll()
	json.NewEncoder(w).Encode(b)
}

func updateBook(w http.ResponseWriter, r *http.Request) {}
