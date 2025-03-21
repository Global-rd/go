package middleware

import (
	"log/slog"
	"net/http"
)

type CorsMiddleware struct {
	logger *slog.Logger
}

func NewCorsMiddleware(logger *slog.Logger) *CorsMiddleware {
	return &CorsMiddleware{
		logger: logger,
	}
}

func (cm *CorsMiddleware) Handle(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cm.logger.Debug("CorsMiddleware.Handle")

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")

		next.ServeHTTP(w, r)
	})
}
