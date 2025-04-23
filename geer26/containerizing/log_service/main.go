package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"main/config"
	"strings"
	"time"

	kafka "github.com/segmentio/kafka-go"
)

type Service struct {
	Reader *kafka.Reader
	Config *config.Cfg
}

type LogMessage struct {
	Log       string    `json:"log"`
	Type      string    `json:"type"`
	Timestamp time.Time `json:"timestamp"`
	ID        string    `json:"id"`
}

func (s *Service) Configure() *Service {
	config, err := config.SetConfig()
	if err != nil {
		slog.Error(err.Error())
		return s
	}
	s.Config = config
	return s
}

func (s *Service) setKafkaReader() *Service {
	brokers := strings.Split(s.Config.Kafkaurl, ",")

	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   brokers,
		Topic:     s.Config.Logtopic,
		Partition: 0,
		MaxBytes:  10e6, // 10MB
		GroupID:   s.Config.Groupid,
	})

	s.Reader = reader

	return s

}

/*
func (l Log) INFO(info string) error {
	logFile, err := os.OpenFile(l.LogfilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		slog.Error("failed to open log file", "error", err)
		return err
	}
	defer logFile.Close()
	slog.New(slog.NewTextHandler(logFile, nil)).Info(info)
	return nil
}
*/

func main() {
	done := make(chan struct{})

	var logservice Service

	logservice.Configure().setKafkaReader()

	go func() {
		defer logservice.Reader.Close()
		log.Println("Log service started...")

		for {
			m, err := logservice.Reader.ReadMessage(context.Background())
			if err != nil {
				log.Fatalf("error at reading queue: %s", err.Error())
			}
			fmt.Printf("Log to save: %s\n", string(m.Value))
		}

	}()

	<-done
}
