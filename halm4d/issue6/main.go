package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"issue6/db"
	"issue6/handler"
	"issue6/middleware"
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

	dbDriver := os.Getenv("DB_DRIVER")
	if dbDriver == "" {
		dbDriver = "mysql"
	}
	dbSource := os.Getenv("DB_SOURCE")
	if dbSource == "" {
		dbSource = "user:password@tcp(localhost:3306)/bookstore"
	}

	// Initialize the database connection
	dbConn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		logger.Error("Error opening database connection", "error", err)
		os.Exit(1)
	}
	defer dbConn.Close()

	mySqlDB := db.NewSQLBookRepository(dbConn, logger)
	err = mySqlDB.Initialize()
	if err != nil {
		logger.Error("Error initializing database", "error", err)
		os.Exit(1)
	}
	bookHandler := handler.NewBookHandler(
		handler.WithLogger(logger),
		handler.WithDatabase(mySqlDB),
	)

	mux.HandleFunc("GET /books", bookHandler.HandleGetAllBooks)
	mux.HandleFunc("GET /books/{id}", bookHandler.HandleGetBookById)
	mux.HandleFunc("POST /books", bookHandler.HandleCreateBook)
	mux.HandleFunc("PUT /books/{id}", bookHandler.HandleUpdateBook)
	mux.HandleFunc("DELETE /books/{id}", bookHandler.HandleDeleteBook)

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
