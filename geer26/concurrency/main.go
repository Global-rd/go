package main

import (
	"log"
	"main/routes"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()
	r.Get("/", routes.BaseHandler)
	log.Println("Server started on port 5000...")
	err := http.ListenAndServe(":5000", r)
	if err != nil {
		log.Fatal("Serving failed")
	}
}
