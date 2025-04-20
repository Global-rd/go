package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
)

func init() {
	logLevel := os.Getenv("LOG_LEVEL")

	var level = slog.LevelDebug
	err := level.UnmarshalText([]byte(logLevel))
	if err != nil {
		slog.Error("Failed to parse log level", "error", err)
	}
	slog.SetLogLoggerLevel(level)
}

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("GET /healthz", func(w http.ResponseWriter, r *http.Request) {
		slog.Debug("Received health check request", "method", r.Method, "url", r.URL.String())
		w.WriteHeader(http.StatusOK)
		_, err := w.Write([]byte("OK"))
		if err != nil {
			http.Error(w, "Error writing response", http.StatusInternalServerError)
			return
		}
	})

	serverPort := os.Getenv("SERVER_PORT")
	if serverPort == "" {
		serverPort = "8080"
	}

	slog.Info("Health checker listening...", "port", serverPort)
	err := http.ListenAndServe(fmt.Sprintf(":%s", serverPort), mux)
	if err != nil {
		slog.Error("Failed to start server", "error", err)
		os.Exit(1)
	}

}
