package main

import (
	"context"
	"log"
	"log/slog"
	"webservice/server"
)

func main() {
	logger := slog.Default()

	srv := server.NewServer(
		nil,
		server.WithLogger(logger.With("name", "server")),
	)
	ctx := context.Background()

	logger.Info("Server started!")

	err := srv.Serve(
		ctx,
		"localhost",
		8080,
	)

	log.Fatal(err)
}
