package task

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
)

func NewRouter(r chi.Router, taskController Controller) chi.Router {
	r.Post("/", func(w http.ResponseWriter, r *http.Request) {
		var task Task

		err := json.NewDecoder(r.Body).Decode(&task)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer r.Body.Close()

		err = taskController.Create(task)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		response, _ := json.Marshal(http.StatusCreated)

		w.Write(response)
	})

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		tasks, err := taskController.Get()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		response, err := json.Marshal(tasks)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(response)
	})

	r.Get("/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")

		task, err := taskController.GetByID(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		response, err := json.Marshal(task)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(response)
	})

	return r
}
