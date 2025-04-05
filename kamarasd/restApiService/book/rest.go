package book

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
)

func NewRouter(r chi.Router, book Controller) chi.Router {
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		books, err := book.HandleBooks()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		response, err := json.Marshal(books)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		w.Write(response)
	})

	r.Get("/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		ret := book.HandleBookById(id)

		writeRet(w, ret)
	})

	r.Delete("/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		ret := book.HandleDeleteBookById(id)

		writeRet(w, ret)
	})

	r.Post("/", func(w http.ResponseWriter, r *http.Request) {
		var newBook Book
		if err := json.NewDecoder(r.Body).Decode(&newBook); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		ret := book.HandleCreateBook(newBook)

		writeRet(w, ret)
	})

	r.Put("/{id}", func(w http.ResponseWriter, r *http.Request) {
		var newBook Book
		id := chi.URLParam(r, "id")
		if err := json.NewDecoder(r.Body).Decode(&newBook); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		ret := book.HandleUpdateBook(newBook, id)

		writeRet(w, ret)
	})

	return r
}

func writeRet(w http.ResponseWriter, ret any) {
	json.NewEncoder(w).Encode(ret)
	w.Header().Set("Content-Type", "application/json")
}
