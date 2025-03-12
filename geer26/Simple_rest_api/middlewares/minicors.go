package middlewares

import (
	"net/http"
)

var AllowedHeaders = ""

func MiniCORS() Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			// BEFORE REQUEST
			if origin := r.Header.Get("Origin"); origin != "" {
				w.Header().Set("Access-Control-Allow-Origin", "*")
				w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
				w.Header().Set("Access-Control-Allow-Headers", AllowedHeaders)
				w.Header().Set("Access-Control-Expose-Headers", "Authorization")
			}

			if r.Method == "OPTIONS" {
				w.WriteHeader(http.StatusOK)
				return
			}
			// Call the next middleware/handler in chain
			f(w, r)
			//AFTER REQUEST, BEFORE DEFERRED
		}
	}
}
