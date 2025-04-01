package routes

import (
	"advrest/logger"
	"advrest/middleware"
	"database/sql"

	"github.com/go-chi/chi/v5"
)

func AttachRoutes(db *sql.DB, log *logger.Log) *chi.Mux {
	r := chi.NewRouter()
	middleware.AttachMiddlewares(r)
	log.INFO("Middlewares attached")

	r.Get("/", Helloroute(db, log))

	r.Mount("/books", BookRoutes(db, log))
	log.INFO("Routes attached")

	return r
}
