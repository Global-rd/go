package server

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"webservice/configs"
)

type OptionFunc func(*Server)

func (f OptionFunc) apply(s *Server) {
	f(s)
}

type Option interface {
	apply(*Server)
}

func WithLogger(logger *slog.Logger) OptionFunc {
	return func(s *Server) {
		s.logger = logger
	}
}

type Server struct {
	logger *slog.Logger
	router http.Handler
}

func NewServer(router http.Handler, opts ...Option) Server {
	server := &Server{
		router: router,
	}

	for _, opt := range opts {
		opt.apply(server)
	}

	return *server
}

func (s Server) Serve(
	ctx context.Context,
	cfg configs.Server,
) error {
	server := http.Server{
		Addr:    fmt.Sprintf("%s:%d", cfg.Address, cfg.Port),
		Handler: s.router,
	}

	s.logger.Info("server started!", "address", cfg.Address, "port", cfg.Port)
	return server.ListenAndServe()
}
