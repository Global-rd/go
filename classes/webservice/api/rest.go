package api

import (
	"fmt"
	"net/http"
	"webservice/book"
	"webservice/container"
	"webservice/task"

	"github.com/go-chi/chi"
)

func NewHttpApi(cont container.Container) *chi.Mux {
	router := chi.NewRouter()
	b := book.NewController(cont)
	t := task.NewController(cont)

	logger := cont.GetLogger().With("name", "middleware")

	router.Use(CreateLogger(logger))

	router.Route("/", func(r chi.Router) {
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprint(w, "hello")
		})

		r.Route("/books", func(r chi.Router) {
			book.NewRouter(r, b)
		})

		r.Route("/tasks", func(r chi.Router) {
			task.NewRouter(r, t)
		})
	})

	return router
}
