package main

import (
	"log/slog"
	"net/http"
	"restApiClient/controller"
	"time"
)

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/books", loggingMiddleware(http.HandlerFunc(controller.HandleBooks)))
	mux.HandleFunc("/books/{id}", loggingMiddleware(http.HandlerFunc(controller.HandleBookById)))

	slog.Info("serves has been started at 8080")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		slog.Error(err.Error())
	}
}

func loggingMiddleware(next http.Handler) func(w http.ResponseWriter, r *http.Request) {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		slog.Info("request", "method", r.Method, "url", r.URL.Path, "time", time.Now())

		next.ServeHTTP(w, r)
	})
}
