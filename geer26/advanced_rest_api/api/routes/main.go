package routes

import (
	"advrest/logger"
	"advrest/middleware"
	"database/sql"

	"github.com/go-chi/chi/v5"
)

//var DbConnection *sql.DB
//var Logger *logger.Log

func AttachRoutes(db *sql.DB, log *logger.Log) *chi.Mux {
	//DbConnection = db
	//Logger = log
	r := chi.NewRouter()
	middleware.AttachMiddlewares(r)
	log.INFO("Middlewares attached")

	r.Get("/", Helloroute(db, log))

	r.Mount("/books", BookRoutes(db, log))
	log.INFO("Routes attached")

	return r
}
