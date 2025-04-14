package api

import (
	"fmt"
	"net/http"
	"webservice/container"

	"github.com/go-chi/chi"
)

func NewHttpApi(cont container.Container) *chi.Mux {
	router := chi.NewRouter()

	logger := cont.GetLogger().With("name", "middleware")

	router.Use(CreateLogger(logger))

	router.Route("/", func(r chi.Router) {
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprint(w, "hello")
		})
	})

	return router
}
