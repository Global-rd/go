package task

import (
	"webservice/container"

	"github.com/google/uuid"
)

type Controller struct {
	cont container.Container
	db   DB
}

func NewController(cont container.Container) Controller {
	db := DB{
		db: cont.GetDB(),
	}

	return Controller{
		cont: cont,
		db:   db,
	}
}

func (c Controller) Create(task Task) error {
	task.ID = uuid.NewString()
	return c.db.Create(task)
}

func (c Controller) GetByID(id string) (Task, error) {
	return Task{}, nil
}

func (c Controller) Get() ([]Task, error) {
	return c.db.Get()
}

func (c Controller) Update(id string) error {
	return nil
}

func (c Controller) Delete(id string) error {
	return nil
}
