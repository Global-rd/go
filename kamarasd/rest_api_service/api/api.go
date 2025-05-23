package api

import (
	book "main/book"
	"main/container"

	"github.com/go-chi/chi"
)

func NewHttpApi(container container.Container) *chi.Mux {
	router := chi.NewRouter()
	bookController := book.NewController(container)

	router.Route("/", func(r chi.Router) {
		r.Route("/books", func(r chi.Router) {
			book.NewRouter(r, bookController)
		})
	})

	return router
}
