package main

import (
	"context"
	"database/sql"
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

	db, err := sql.Open("postgres", cfg.DB.ConnectionString())
	if err != nil {
		logger.Error("couldn't connect to db", slog.String("err", err.Error()))
		os.Exit(1)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		logger.Error("couldn't ping the db", slog.String("err", err.Error()))
		os.Exit(1)
	}

	cont := container.NewContainer(
		logger,
		db,
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
