package payment

import (
	"context"
	"encoding/json"
	"log/slog"
	"webservice/container"

	kafka "github.com/segmentio/kafka-go"
)

type Kafka struct {
	writer *kafka.Writer
	logger *slog.Logger
}

func NewKafka(cont container.Container) Kafka {
	return Kafka{
		writer: cont.Kafka(),
		logger: cont.GetLogger(),
	}
}

func (k Kafka) WritePayment(ctx context.Context, model Payment) error {
	value, err := json.Marshal(model)
	if err != nil {
		return err
	}

	// time out pattern with contex done

	err = k.writer.WriteMessages(ctx,
		kafka.Message{
			Key:   []byte("Key-A"),
			Value: value,
		},
	)
	if err != nil {
		k.logger.Error(err.Error())
	}

	return err
}
