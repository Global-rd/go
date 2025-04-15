package server

import (
	"context"
	"fmt"
	"log/slog"
	configs "main/configs"
	"net/http"
)

type OptionalFunc func(*Server)

type Option interface {
	apply(*Server)
}

type Server struct {
	router http.Handler
}

func NewServer(router http.Handler, options ...Option) Server {
	server := &Server{
		router: router,
	}

	for _, option := range options {
		option.apply(server)
	}

	return *server
}

func (s Server) StartServer(context context.Context, config configs.Server) error {
	server := http.Server{
		Addr:    fmt.Sprintf("%s:%d", config.Address, config.Port),
		Handler: s.router,
	}

	slog.Info("Server start", "address", config.Address, "port", config.Port)
	return server.ListenAndServe()

}
