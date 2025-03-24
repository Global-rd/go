package service

import (
	"advrest/config"
	"advrest/logger"
	"advrest/routes"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	_ "github.com/lib/pq"
)

type Service struct {
	Router    *chi.Mux
	Server    *http.Server
	Config    *config.Cfg
	Db        *sql.DB
	Logger    *logger.Log
	InitError error
}

func ServiceBuilder() *Service {
	service := Service{
		Server: &http.Server{},
	}
	return &service
}

func (s *Service) CreateLogger(option ...logger.Option) *Service {
	if s.InitError != nil {
		return s
	}
	logger, err := logger.InitLogger(option...)
	if err != nil {
		s.InitError = err
		return s
	}

	s.Logger = logger
	s.Logger.INFO("Logger init")
	return s
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

	db, err := sql.Open("postgres", config.BuildConnectionString(s.Config.DB))
	if err != nil {
		s.InitError = err
		s.Logger.ERROR(fmt.Sprintf("Database connection failed: %s", err))
		return s
	}
	//defer db.Close()

	err = db.Ping()
	if err != nil {
		s.InitError = err
		s.Logger.ERROR(fmt.Sprintf("Database ping failed: %s", err))
		return s
	}

	s.Db = db
	s.Logger.INFO("Database connected")
	return s
}

func (s *Service) AttachRoutes() *Service {
	if s.InitError != nil {
		return s
	}
	s.Server.Handler = routes.AttachRoutes(s.Db, s.Logger)
	s.Logger.INFO("Routes attached")
	return s
}

func (s *Service) Run() (*Service, error) {
	if s.InitError != nil {
		s.Logger.ERROR(fmt.Sprintf("Service startup failed: %s", s.InitError))
		return s, s.InitError
	}

	s.Server.Addr = fmt.Sprintf(":%d", s.Config.Server.Port)
	s.Server.ReadTimeout = time.Duration(s.Config.Server.ReadTimeout * int(time.Second))
	s.Server.WriteTimeout = time.Duration(s.Config.Server.WriteWimeout * int(time.Second))

	log.Printf("\nService started up on port %s...", s.Server.Addr)
	s.Logger.INFO(fmt.Sprintf("Service started up on port %s", s.Server.Addr))
	err := s.Server.ListenAndServe()
	if err != nil {
		s.Logger.ERROR(fmt.Sprintf("Service listen failed: %s", err))
		return s, err
	}
	return s, nil
}
