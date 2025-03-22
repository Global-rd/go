package routes

import (
	"net/http"

	"github.com/go-chi/chi"
)

type BookController struct {
}

func (b BookController) ListBooks(w http.ResponseWriter, r *http.Request)  {}
func (b BookController) GetBooks(w http.ResponseWriter, r *http.Request)   {}
func (b BookController) CreateBook(w http.ResponseWriter, r *http.Request) {}
func (b BookController) UpdateBook(w http.ResponseWriter, r *http.Request) {}
func (b BookController) DeleteBook(w http.ResponseWriter, r *http.Request) {}

func BookRoutes() chi.Router {
	r := chi.NewRouter()
	bookHandler := BookController{}
	r.Get("/", bookHandler.ListBooks)
	r.Post("/", bookHandler.CreateBook)
	r.Get("/{id}", bookHandler.GetBooks)
	r.Put("/{id}", bookHandler.UpdateBook)
	r.Delete("/{id}", bookHandler.DeleteBook)
	return r
}
