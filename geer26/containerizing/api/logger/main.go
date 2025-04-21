package logger

import (
	"advrest/config"
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/google/uuid"
	kafka "github.com/segmentio/kafka-go"
)

type LogMessage struct {
	Log       string    `json:"log"`
	Type      string    `json:"type"`
	Timestamp time.Time `json:"timestamp"`
	ID        uuid.UUID `json:"id"`
}

type Log struct {
	KafkaWriter *kafka.Writer
}

func newKafkaWriter(kafkaURL, topic string) *kafka.Writer {
	return &kafka.Writer{
		Addr:                   kafka.TCP(kafkaURL),
		Topic:                  topic,
		Balancer:               &kafka.LeastBytes{},
		AllowAutoTopicCreation: true,
	}
}

func (l Log) INFO(info string) error {
	message := LogMessage{
		ID:        uuid.New(),
		Log:       info,
		Type:      "INFO",
		Timestamp: time.Now(),
	}

	bytearray, err := json.Marshal(message)
	if err != nil {
		return err
	}

	msg := kafka.Message{
		Key:   []byte(fmt.Sprint(uuid.New())),
		Value: bytearray,
	}

	err = l.KafkaWriter.WriteMessages(context.Background(), msg)
	if err != nil {
		return err
	}
	return nil
}

func (l Log) WARNING(info string) error {
	message := LogMessage{
		ID:        uuid.New(),
		Log:       info,
		Type:      "WARNING",
		Timestamp: time.Now(),
	}

	bytearray, err := json.Marshal(message)
	if err != nil {
		return err
	}

	msg := kafka.Message{
		Key:   []byte(fmt.Sprint(uuid.New())),
		Value: bytearray,
	}

	err = l.KafkaWriter.WriteMessages(context.Background(), msg)
	if err != nil {
		return err
	}
	return nil
}

func (l Log) ERROR(info string) error {
	message := LogMessage{
		ID:        uuid.New(),
		Log:       info,
		Type:      "ERROR",
		Timestamp: time.Now(),
	}

	bytearray, err := json.Marshal(message)
	if err != nil {
		return err
	}

	msg := kafka.Message{
		Key:   []byte(fmt.Sprint(uuid.New())),
		Value: bytearray,
	}

	err = l.KafkaWriter.WriteMessages(context.Background(), msg)
	if err != nil {
		return err
	}
	return nil
}

func InitLogger(config config.Server) (*Log, error) {

	writer := newKafkaWriter(config.KAFKAURL, config.LOGTOPIC)
	//defer writer.Close()
	return &Log{KafkaWriter: writer}, nil
}
