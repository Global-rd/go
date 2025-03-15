package middlewares

import (
	"fmt"
	"log/slog"
	"net/http"
	"time"
)

func Logger() Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			// BEFORE REQUEST
			start := time.Now()
			var method = r.Method
			var path = r.URL.Path

			// Call the next middleware/handler in chain
			f(w, r)
			//AFTER REQUEST, BEFORE DEFERRED
			slog.Info(fmt.Sprintf("Time: %s, Method: %s, Path: %s, Elapsed time: %s", start, method, path, time.Since(start)))
		}
	}
}
