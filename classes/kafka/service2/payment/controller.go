package payment

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

func (c Controller) Create(payment Payment) error {
	payment.ID = uuid.NewString()
	return c.db.Create(payment)
}

func (c Controller) GetByID(id string) (Payment, error) {
	return Payment{}, nil
}

func (c Controller) Get() ([]Payment, error) {
	return c.db.Get()
}

func (c Controller) Update(id string) error {
	return nil
}

func (c Controller) Delete(id string) error {
	return nil
}
