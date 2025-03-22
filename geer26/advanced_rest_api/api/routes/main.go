package routes

import (
	"advrest/middleware"
	"database/sql"

	"github.com/go-chi/chi/v5"
)

var Database *sql.DB

func AttachRoutes(db *sql.DB) *chi.Mux {
	Database = db
	r := chi.NewRouter()
	middleware.AttachMiddlewares(r)

	r.Get("/", Helloroute)

	r.Mount("/books", BookRoutes())

	return r
}
