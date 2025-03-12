package routes

import (
	"encoding/json"
	"main/database"
	"net/http"
	"strconv"
)

func GetBooks(w http.ResponseWriter, r *http.Request) {
	book_id := r.PathValue("id")

	books, err := database.DialStore()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	if book_id == "" {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(books)
		return
	}
	id, err := strconv.Atoi(book_id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	for _, book := range books.Books {
		if book.Id == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(book)
			return
		}
	}
}
