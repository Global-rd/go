package routes

import (
	"encoding/json"
	"main/database"
	"net/http"

	"github.com/google/uuid"
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

func addBook(w http.ResponseWriter, r *http.Request) {
	new_id := uuid.New().String()
	var new_book database.Book
	err := json.NewDecoder(r.Body).Decode(&new_book)
	if err != nil {
		http.Error(w, "invalid json body", http.StatusBadRequest)
	}
	new_book.Id = string(new_id)
	books, err := database.DialStore()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_, err = books.CreateOne(new_book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := books.LoadStore(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	b, _ := books.FindAll()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(b)
}

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
	_, err = books.DeleteOne("Id", book_id)
	if err != nil {
		http.Error(w, "book not found", http.StatusNotFound)
		return
	}
	if err := books.LoadStore(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	b, _ := books.FindAll()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(b)
}

func updateBook(w http.ResponseWriter, r *http.Request) {
	var modified_book database.Book
	err := json.NewDecoder(r.Body).Decode(&modified_book)
	if err != nil {
		http.Error(w, "invalid json body", http.StatusBadRequest)
		return
	}
	old_id := modified_book.Id
	books, err := database.DialStore()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_, err = books.UpdateOne(old_id, modified_book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	b, _ := books.FindAll()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(b)
}
