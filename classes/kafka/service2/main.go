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
	"webservice/payment"
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

	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:     []string{cfg.Kafka.Address},
		Topic:       "payment",
		Partition:   0,
		MaxBytes:    10e6, // 10MB
		Logger:      kafka.LoggerFunc(func(s string, i ...interface{}) { logger.Info(s, i...) }),
		ErrorLogger: kafka.LoggerFunc(func(s string, i ...interface{}) { logger.Error(s, i...) }),
		GroupID:     cfg.Kafka.GroupID,
	})

	cont := container.NewContainer(
		logger,
		db,
		r,
	)

	srv := server.NewServer(
		api.NewHttpApi(cont),
		server.WithLogger(logger.With("name", "server")),
	)

	ctx := context.Background()

	go func() {
		kafka := payment.NewKafka(cont)
		err := kafka.Loop(ctx)
		if err != nil {
			logger.Error(err.Error())
		}
	}()

	logger.Info("server constracted!")

	err = srv.Serve(
		ctx,
		cfg.Server,
	)

	log.Fatal(err)
}
