package middlewares

import (
	"log"
	"net/http"
	"time"
)

func Logger() Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			// BEFORE REQUEST
			start := time.Now()
			defer func() {
				log.Printf("Method: %s, Path: %s, elapsed time: %s", r.Method, r.URL.Path, time.Since(start))
			}()

			// Call the next middleware/handler in chain
			f(w, r)
			//AFTER REQUEST, BEFORE DEFERRED
		}
	}
}
