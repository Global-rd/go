package main

import (
	"context"
	"log"
	"log/slog"
	"os"
	"webservice/api"
	"webservice/configs"
	"webservice/container"
	"webservice/server"

	_ "github.com/lib/pq"
	"github.com/segmentio/kafka-go"
)

func main() {
	logger := slog.Default()

	cfg, err := configs.Parse()
	if err != nil {
		logger.Error("parsing config file failed", "err", err.Error())
		os.Exit(1)
	}

	w := &kafka.Writer{
		Addr:                   kafka.TCP(cfg.Kafka.Address),
		Topic:                  "payment",
		Balancer:               &kafka.LeastBytes{},
		AllowAutoTopicCreation: true,
	}

	cont := container.NewContainer(
		logger,
		w,
	)

	srv := server.NewServer(
		api.NewHttpApi(cont),
		server.WithLogger(logger.With("name", "server")),
	)

	ctx := context.Background()

	logger.Info("server constracted!")

	err = srv.Serve(
		ctx,
		cfg.Server,
	)

	log.Fatal(err)
}
