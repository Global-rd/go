package main

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"os"
	"webservice-std/db"
)

func main() {
	err := db.LoadCache()
	if err != nil {
		slog.Error("DB Cache initialization failed",
			"Error: ", err.Error())
		os.Exit(1)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/books", handleGetAllBooks)
	mux.HandleFunc("/books/{id}", handleGetBookById)

	originsWrapper := allowOriginsMiddleware(mux)
	loggingWrapper := loggingMiddleware(originsWrapper)

	err = http.ListenAndServe(":8080", loggingWrapper)
	if err != nil {
		slog.Error(err.Error())
	}
}

func allowOriginsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		slog.Info("Received request with",
			"path", r.URL.Path,
			"method", r.Method,
			"headers", r.Header,
		)
		next.ServeHTTP(w, r)
	})
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
