package routes

import (
	"advrest/middleware"

	"github.com/go-chi/chi/v5"
)

func AttachRoutes() *chi.Mux {
	r := chi.NewRouter()
	middleware.AttachMiddlewares(r)

	r.Get("/", Helloroute)

	r.Mount("/books", BookRoutes())

	return r
}
