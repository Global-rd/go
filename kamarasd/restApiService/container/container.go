package container

import "database/sql"

type Container struct {
	db *sql.DB
}

func NewContainer(db *sql.DB) Container {
	return Container{
		db: db,
	}
}

func (container Container) GetDB() *sql.DB {
	return container.db
}
