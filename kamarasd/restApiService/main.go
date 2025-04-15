package main

import (
	"context"
	"database/sql"
	"log/slog"
	"main/api"
	"main/container"
	"main/server"
	"net/http"
	"os"
	"time"

	config "main/configs"

	_ "github.com/lib/pq"
)

func main() {

	logger := slog.Default()

	cfg, err := config.ReadConfig()
	if err != nil {
		logger.Error("parsing config file failed", "err", err.Error())
		os.Exit(1)
	}

	database, err := sql.Open("postgres", cfg.DB.ConnectionString())
	if err != nil {
		logger.Error("couldn't connect to db", slog.String("err", err.Error()))
		os.Exit(1)
	}
	defer database.Close()

	err = database.Ping()
	if err != nil {
		logger.Error("couldn't ping the db", slog.String("err", err.Error()))
		os.Exit(1)
	}

	cont := container.NewContainer(
		database,
	)

	serv := server.NewServer(
		api.NewHttpApi(cont),
	)

	contxt := context.Background()

	err = serv.StartServer(
		contxt,
		cfg.Server,
	)

	if err != nil {
		slog.Error(err.Error())
	}
}

func LoggingMiddleware(next http.Handler) func(w http.ResponseWriter, r *http.Request) {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		slog.Info("request", "method", r.Method, "url", r.URL.Path, "time", time.Now())

		next.ServeHTTP(w, r)
	})
}
