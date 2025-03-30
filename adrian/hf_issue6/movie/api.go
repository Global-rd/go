package movie

import "github.com/go-chi/chi"

func NewMovieRouter(r chi.Router, controller *MovieController) chi.Router {
	r.Get("/{id}", controller.GetMovie)
	r.Post("/", controller.CreateMovie)
	r.Put("/{id}", controller.UpdateMovie)
	r.Delete("/{id}", controller.DeleteMovie)

	return r
}
