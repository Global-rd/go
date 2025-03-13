package api

import (
	"fmt"
	"log/slog"
	"net/http"
	"webservice/book"

	"github.com/go-chi/chi"
)

func NewHttpApi(logger *slog.Logger) *chi.Mux {
	router := chi.NewRouter()
	b := book.NewController()

	router.Use(CreateLogger(logger))

	router.Route("/", func(r chi.Router) {
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprint(w, "hello")
		})

		r.Route("/books", func(r chi.Router) {
			book.NewRouter(r, b)
		})
	})

	return router
}
