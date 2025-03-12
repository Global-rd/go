package main

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"webservice-std/db"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/books", handleGetAllBooks)
	mux.HandleFunc("/books/{id}", handleGetBookById)

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		slog.Error(err.Error())
	}
}

func handleGetAllBooks(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(db.Cache)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	return
}

func handleGetBookById(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
	id := r.PathValue("id")
	book, err := db.GetBook(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(book)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	return
}
