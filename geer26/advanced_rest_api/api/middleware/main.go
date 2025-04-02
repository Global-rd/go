package middleware

import (
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

func AttachMiddlewares(r *chi.Mux) error {
	r.Use(
		middleware.Logger,
	)
	return nil
}
