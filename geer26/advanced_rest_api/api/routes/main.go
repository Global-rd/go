package routes

import (
	"advrest/logger"
	"advrest/middleware"
	"database/sql"

	"github.com/go-chi/chi/v5"
)

var DbConnection *sql.DB
var Logger *logger.Log

func AttachRoutes(db *sql.DB, log *logger.Log) *chi.Mux {
	DbConnection = db
	Logger = log
	r := chi.NewRouter()
	middleware.AttachMiddlewares(r)
	Logger.INFO("Middlewares attached")

	r.Get("/", Helloroute)

	r.Mount("/books", BookRoutes())
	Logger.INFO("Routes attached")

	return r
}
