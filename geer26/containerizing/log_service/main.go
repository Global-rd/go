package main

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	kafka "github.com/segmentio/kafka-go"
)

type LogMessage struct {
	Log       string    `json:"log"`
	Type      string    `json:"type"`
	Timestamp time.Time `json:"timestamp"`
	ID        string    `json:"id"`
}

const (
	KAFKAURL = "localhost:9092"
	TOPIC    = "log_service"
)

func getKafkaReader(kafkaURL, topic, groupID string) *kafka.Reader {
	brokers := strings.Split(kafkaURL, ",")

	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   brokers,
		Topic:     topic,
		Partition: 0,
		MaxBytes:  10e6, // 10MB
		GroupID:   groupID,
	})

	return reader

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

	go func() {
		reader := getKafkaReader(KAFKAURL, TOPIC, "logger_service")
		defer reader.Close()

		for {
			m, err := reader.ReadMessage(context.Background())
			if err != nil {
				log.Fatalf("error at reading queue: %s", err.Error())
			}
			fmt.Printf("message: %s\n", string(m.Value))
		}

	}()

	<-done
}
