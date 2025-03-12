package main

import "log/slog"

const (
	DefaultServerName string = "finance"
)

type HTTPServer struct {
	addr    string
	port    int
	name    string
	logger  *slog.Logger
	timeout int
}

type Option interface {
	apply(*HTTPServer)
}

type ServerConfig struct {
	timeout int
	name    string
}

func (sc ServerConfig) constructServerName() string {
	// ...
	return sc.name + "constructed string"
}

func (sc ServerConfig) apply(server *HTTPServer) {
	server.name = sc.constructServerName()
	server.timeout = sc.timeout
}

type OptionFunc func(*HTTPServer)

func (f OptionFunc) apply(server *HTTPServer) {
	f(server)
}

func WithLogger(logger *slog.Logger) OptionFunc {
	return OptionFunc(func(h *HTTPServer) {
		h.logger = logger
	})
}

func WithName(name string) OptionFunc {
	return OptionFunc(func(h *HTTPServer) {
		h.name = name
	})
}

func NewHTTPServer(addr string, port int, options ...Option) HTTPServer {
	server := &HTTPServer{
		addr: addr,
		port: port,
		name: DefaultServerName,
	}

	for _, option := range options {
		option.apply(server)
	}

	return *server
}
