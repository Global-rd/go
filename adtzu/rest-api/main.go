package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"rest-api/database"
	"rest-api/routes"
)

func main() {
	db := database.NewMemDB()
	router := mux.NewRouter()
	routes.SetupBookRoutes(router, db)
	log.Fatal(http.ListenAndServe(":8000", router))
}
