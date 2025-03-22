package service

import (
	"advrest/config"
	"advrest/routes"
	"database/sql"
	"log"
	"net/http"

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
		Server: &http.Server{
			Addr: ":5000",
		},
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
	log.Println("Service started up...")
	err := s.Server.ListenAndServe()
	if err != nil {
		return err
	}
	return nil
}
