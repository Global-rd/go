package service

import (
	"advrest/config"
	"advrest/routes"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
)

type Service struct {
	Router    *chi.Mux
	Server    *http.Server
	Config    *config.Cfg
	Db        *sql.DB
	InitError error
}

func ServiceBuilder() *Service {
	service := Service{
		Server: &http.Server{},
	}
	return &service
}

func (s *Service) Configure() *Service {
	if s.InitError != nil {
		return s
	}
	config, err := config.SetConfig()
	if err != nil {
		s.InitError = err
		return s
	}
	s.Config = config
	return s
}

func (s *Service) Connect() *Service {
	if s.InitError != nil {
		return s
	}
	return s
}

func (s *Service) AttachRoutes() *Service {
	if s.InitError != nil {
		return s
	}
	s.Server.Handler = routes.AttachRoutes()
	return s
}

func (s *Service) Run() error {
	if s.InitError != nil {
		return s.InitError
	}

	s.Server.Addr = fmt.Sprintf(":%d", s.Config.Server.Port)
	s.Server.ReadTimeout = time.Duration(s.Config.Server.ReadTimeout * int(time.Second))
	s.Server.WriteTimeout = time.Duration(s.Config.Server.WriteWimeout * int(time.Second))

	log.Printf("Service started up on port %s...", s.Server.Addr)
	err := s.Server.ListenAndServe()
	if err != nil {
		return err
	}
	return nil
}
