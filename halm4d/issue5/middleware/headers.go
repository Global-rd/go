package middleware

import (
	"github.com/google/uuid"
	"log/slog"
	"net/http"
)

type HeadersMiddleware struct {
	logger *slog.Logger
}

func NewHeadersMiddleware(logger *slog.Logger) *HeadersMiddleware {
	return &HeadersMiddleware{
		logger: logger,
	}
}

func (hm *HeadersMiddleware) Handle(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		hm.logger.Debug("HeadersMiddleware.Handle")

		cid := r.Header.Get(CORRELATION_ID_HEADER)
		if cid == "" {
			cid = uuid.NewString()
		}
		r.Header.Add(CORRELATION_ID_HEADER, cid)
		w.Header().Add(CORRELATION_ID_HEADER, cid)

		next.ServeHTTP(w, r)
	})
}

const CORRELATION_ID_HEADER = "X-Correlation-Id"
