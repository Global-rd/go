package api

import (
	"log/slog"
	"net/http"
	"time"
)

func CreateLogger(logger *slog.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {

		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			t := time.Now()
			logger.Info("request", "method", r.Method, "URL", r.URL, "time", t)
			next.ServeHTTP(w, r)
		})
	}
}
