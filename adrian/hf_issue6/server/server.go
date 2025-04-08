package server

import (
	"errors"
	"fmt"
	"full-webservice/config"
	"log/slog"
	"net/http"
)

type Server struct {
	logger *slog.Logger
	router http.Handler
	addr   string
	port   int
}

type Option func(*Server)

func NewServer(router http.Handler, cfg config.ServerCfg, opts ...Option) (*Server, error) {

	if cfg.Port == 0 || cfg.Address == "" {
		return nil, errors.New("invalid server configuration")
	}

	server := &Server{
		logger: slog.Default(),
		router: router,
		addr:   cfg.Address,
		port:   cfg.Port,
	}

	for _, opt := range opts {
		opt(server)
	}

	return server, nil
}

func WithLogger(logger *slog.Logger) Option {
	return func(s *Server) {
		s.logger = logger
	}
}

func (s *Server) Start() error {
	httpServer := http.Server{
		Addr:    fmt.Sprintf("%s:%d", s.addr, s.port),
		Handler: s.router,
	}
	s.logger.Info("starting server", slog.String("addr", s.addr), slog.Int("port", s.port))
	return httpServer.ListenAndServe()
}
