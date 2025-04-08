package payment

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
)

func NewRouter(r chi.Router, paymentController Controller) chi.Router {
	r.Post("/", func(w http.ResponseWriter, r *http.Request) {
		var payment Payment

		err := json.NewDecoder(r.Body).Decode(&payment)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer r.Body.Close()

		err = paymentController.Create(payment)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		response, _ := json.Marshal(http.StatusCreated)

		w.Write(response)
	})

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		payments, err := paymentController.Get()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		response, err := json.Marshal(payments)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(response)
	})

	r.Get("/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")

		payment, err := paymentController.GetByID(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		response, err := json.Marshal(payment)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(response)
	})

	return r
}
