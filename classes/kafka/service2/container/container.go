package container

import (
	"database/sql"
	"log/slog"

	kafka "github.com/segmentio/kafka-go"
)

type Container struct {
	logger *slog.Logger
	db     *sql.DB
	kafka  *kafka.Reader
}

func NewContainer(
	logger *slog.Logger,
	db *sql.DB,
	kafka *kafka.Reader,
) Container {
	return Container{
		logger: logger,
		db:     db,
		kafka:  kafka,
	}
}

func (c Container) GetLogger() *slog.Logger {
	return c.logger
}

func (c Container) GetDB() *sql.DB {
	return c.db
}

func (c Container) Kafka() *kafka.Reader {
	return c.kafka
}
