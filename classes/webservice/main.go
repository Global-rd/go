package main

import (
	"context"
	"log"
	"log/slog"
	"webservice/api"
	"webservice/server"
)

func main() {
	logger := slog.Default()

	srv := server.NewServer(
		api.NewHttpApi(logger.With("name", "middleware")),
		server.WithLogger(logger.With("name", "server")),
	)

	ctx := context.Background()

	logger.Info("Server constracted!")

	err := srv.Serve(
		ctx,
		"localhost",
		8080,
	)

	log.Fatal(err)
}
