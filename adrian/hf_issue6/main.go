package main

import (
	"database/sql"
	"fmt"
	"full-webservice/config"
	"full-webservice/movie"
	"full-webservice/server"
	"github.com/go-chi/chi"
	_ "github.com/lib/pq"
	"log/slog"
	"os"
)

func connectDb(cfg config.DbCfg) (*sql.DB, error) {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.DbHost, cfg.DbPort, cfg.DbUser, cfg.DbPass, cfg.DbName)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	return db, nil

}

func main() {
	logger := slog.Default()
	cfg, err := config.LoadConfig()
	if err != nil {
		logger.Error("failed to read configuration", slog.String("err", err.Error()))
		os.Exit(1)
	}
	var dbConn *sql.DB
	dbConn, err = connectDb(cfg.Db)
	if err != nil {
		logger.Error("failed to connect to database", slog.String("err", err.Error()))
		os.Exit(1)
	}
	defer dbConn.Close()

	controller := movie.NewMovieController(dbConn)
	router := chi.NewRouter()
	router.Route("/movies", func(r chi.Router) {
		movie.NewMovieRouter(r, controller)
	})

	var srv *server.Server
	srv, err = server.NewServer(router, cfg.Server, server.WithLogger(logger))
	if err != nil {
		logger.Error("failed to create server", slog.String("err", err.Error()))
		os.Exit(1)
	}
	err = srv.Start()
	if err != nil {
		logger.Error("failed to start server", slog.String("err", err.Error()))
		os.Exit(1)
	}
}
