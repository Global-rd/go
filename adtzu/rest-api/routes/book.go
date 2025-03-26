package routes

import (
	"github.com/gorilla/mux"

	"rest-api/database"
	"rest-api/handlers"
)

func SetupBookRoutes(router *mux.Router, db *database.MemDB) {
	handler := handlers.NewBookHandler(db)
	bookRouter := router.PathPrefix("/books").Subrouter()
	bookRouter.HandleFunc("", handler.CreateBook).Methods("POST")
	bookRouter.HandleFunc("/{id}", handler.GetBook).Methods("GET")
	bookRouter.HandleFunc("/{id}", handler.UpdateBook).Methods("PUT")
	bookRouter.HandleFunc("/{id}", handler.DeleteBook).Methods("DELETE")
}
