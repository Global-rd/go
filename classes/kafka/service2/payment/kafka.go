package payment

import (
	"context"
	"encoding/json"
	"log/slog"
	"webservice/container"

	"github.com/segmentio/kafka-go"
)

type Kafka struct {
	reader *kafka.Reader
	logger *slog.Logger
	db     DB
}

func NewKafka(cont container.Container) Kafka {
	db := DB{
		db: cont.GetDB(),
	}

	return Kafka{
		reader: cont.Kafka(),
		logger: cont.GetLogger(),
		db:     db,
	}
}

func (k Kafka) Loop(ctx context.Context) error {
	for {
		m, err := k.reader.ReadMessage(ctx)
		if err != nil {
			k.logger.Error(err.Error())
			break
		}

		var p Payment

		json.Unmarshal(m.Value, &p)

		err = k.db.Create(p)

		if err != nil {
			k.logger.Error(err.Error())
		}
	}

	if err := k.reader.Close(); err != nil {
		k.logger.Error(err.Error())
		return err
	}

	return nil
}
