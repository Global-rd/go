package routes

import (
	"advrest/middleware"
	"database/sql"

	"github.com/go-chi/chi/v5"
)

var DbConnection *sql.DB

func AttachRoutes(db *sql.DB) *chi.Mux {
	DbConnection = db
	r := chi.NewRouter()
	middleware.AttachMiddlewares(r)

	r.Get("/", Helloroute)

	r.Mount("/books", BookRoutes())

	return r
}
