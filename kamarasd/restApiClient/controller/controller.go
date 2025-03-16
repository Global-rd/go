package controller

import (
	"encoding/json"
	"log/slog"
	"net/http"
	book "restApiClient/models"
	jsonService "restApiClient/services"
	"strconv"
)

func HandleBooks(w http.ResponseWriter, r *http.Request) {
	books := book.NewBookModel()
	jsonService.OpenJsonFile(books, "books.json")

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(books)
	if err != nil {
		slog.Error(err.Error())
	}

}

func HandleBookById(w http.ResponseWriter, r *http.Request) {
	books := book.NewBookModel()
	jsonService.OpenJsonFile(books, "books.json")

	w.Header().Set("Content-Type", "application/json")
	bookIdValue := r.PathValue("id")
	bookId, err := strconv.Atoi(bookIdValue)

	if err != nil {
		slog.Error(err.Error())
	}

	flag := false

	for _, book := range books.Books {
		if book.ID == bookId {
			json.NewEncoder(w).Encode(book)
			flag = true
		}
	}

	if !flag {
		json.NewEncoder(w).Encode("No book selected")
	}
}
