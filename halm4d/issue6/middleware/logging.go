package middleware

import (
	"bytes"
	"log/slog"
	"net/http"
)

type LoggingMiddleware struct {
	logger *slog.Logger
}

func NewLoggingMiddleware(logger *slog.Logger) *LoggingMiddleware {
	return &LoggingMiddleware{
		logger: logger,
	}
}

type responseCaptureWriter struct {
	http.ResponseWriter
	statusCode int
	body       *bytes.Buffer
}

func newResponseCaptureWriter(w http.ResponseWriter) *responseCaptureWriter {
	return &responseCaptureWriter{
		ResponseWriter: w,
		statusCode:     http.StatusOK,
		body:           &bytes.Buffer{},
	}
}

func (w *responseCaptureWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func (w *responseCaptureWriter) WriteHeader(statusCode int) {
	w.statusCode = statusCode
	w.ResponseWriter.WriteHeader(statusCode)
}

func (lm *LoggingMiddleware) Handle(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		lm.logger.Debug("LoggingMiddleware.Handle")

		cid := r.Header.Get(CORRELATION_ID_HEADER)
		path := r.URL.Path
		method := r.Method

		lm.logger.Info("--> REQUEST",
			"method", method,
			"url", path,
			"correlation-id", cid,
			"headers", r.Header,
		)
		lm.logger.Debug("--> REQUEST BODY",
			"method", method,
			"url", path,
			"correlation-id", cid,
			"body", r.Body,
		)

		captureWriter := newResponseCaptureWriter(w)
		next.ServeHTTP(captureWriter, r)

		lm.logger.Info(
			"<-- RESPONSE",
			"status", captureWriter.statusCode,
			"method", method,
			"url", path,
			"correlation-id", cid,
			"headers", w.Header(),
		)
		lm.logger.Debug(
			"<-- RESPONSE BODY",
			"status", captureWriter.statusCode,
			"method", method,
			"url", path,
			"correlation-id", cid,
			"body", captureWriter.body.String(),
		)
	})
}
