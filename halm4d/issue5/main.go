package main

import (
	"fmt"
	"issue5/db"
	"issue5/handler"
	"issue5/middleware"
	"log/slog"
	"net/http"
	"os"
)

func main() {

	logLevelEnv := os.Getenv("LOG_LEVEL")
	if logLevelEnv == "" {
		logLevelEnv = "DEBUG"
	}

	var level slog.Level
	err := level.UnmarshalText([]byte(logLevelEnv))
	if err != nil {
		slog.Error("Error unmarshalling log level", "error", err)
		os.Exit(1)
	}
	logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level: level,
	}))

	mux := http.NewServeMux()

	inMemoryDB := db.NewInMemoryBookRepository()
	bookHandler := handler.NewBookHandler(
		handler.WithLogger(logger),
		handler.WithDatabase(inMemoryDB),
	)

	mux.HandleFunc("/books", bookHandler.HandleBooks)
	mux.HandleFunc("/books/{id}", bookHandler.HandleBook)

	middlewareChain := middleware.NewChain()
	middlewareChain.Use(middleware.NewHeadersMiddleware(logger))
	middlewareChain.Use(middleware.NewCorsMiddleware(logger))
	middlewareChain.Use(middleware.NewLoggingMiddleware(logger))
	wrappedMux := middlewareChain.Handle(mux)

	serverPort := os.Getenv("SERVER_PORT")
	if serverPort == "" {
		serverPort = ":8080"
	}

	logger.Info(fmt.Sprintf("Starting server on %s", serverPort))
	if err := http.ListenAndServe(serverPort, wrappedMux); err != nil {
		logger.Error("Error starting server", "error", err)
	}
}
