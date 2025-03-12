package routes

import (
	"log"
	"net/http"
)

func GetBooks(w http.ResponseWriter, r *http.Request) {
	book_id := r.PathValue("id")
	log.Println("PATH_VALUE: ", book_id)
}
