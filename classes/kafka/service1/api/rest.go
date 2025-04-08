package api

import (
	"fmt"
	"net/http"
	"webservice/container"
	"webservice/payment"

	"github.com/go-chi/chi"
)

func NewHttpApi(cont container.Container) *chi.Mux {
	router := chi.NewRouter()
	p := payment.NewController(cont)

	logger := cont.GetLogger().With("name", "middleware")

	router.Use(CreateLogger(logger))

	router.Route("/", func(r chi.Router) {
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprint(w, "hello")
		})

		r.Route("/payment", func(r chi.Router) {
			payment.NewRouter(r, p)
		})
	})

	return router
}
