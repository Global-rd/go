package container

import (
	"database/sql"
	"log/slog"
)

type Container struct {
	logger *slog.Logger
	db     *sql.DB
}

func NewContainer(
	logger *slog.Logger,
	db *sql.DB,
) Container {
	return Container{
		logger: logger,
		db:     db,
	}
}

func (c Container) GetLogger() *slog.Logger {
	return c.logger
}

func (c Container) GetDB() *sql.DB {
	return c.db
}
