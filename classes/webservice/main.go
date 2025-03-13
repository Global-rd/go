package main

import (
	"context"
	"log"
	"log/slog"
	"os"
	"webservice/api"
	"webservice/configs"
	"webservice/server"
)

func main() {
	logger := slog.Default()

	cfg, err := configs.Parse()
	if err != nil {
		logger.Error("parsing config file failed", "err", err.Error())
		os.Exit(1)
	}

	srv := server.NewServer(
		api.NewHttpApi(logger.With("name", "middleware")),
		server.WithLogger(logger.With("name", "server")),
	)

	ctx := context.Background()

	logger.Info("server constracted!")

	err = srv.Serve(
		ctx,
		cfg.Address,
		cfg.Port,
	)

	log.Fatal(err)
}
