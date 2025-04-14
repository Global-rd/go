package payment

import (
	"context"
	"webservice/container"

	"github.com/google/uuid"
)

type Controller struct {
	cont  container.Container
	kafka Kafka
}

func NewController(cont container.Container) Controller {
	kafka := NewKafka(cont)

	return Controller{
		cont:  cont,
		kafka: kafka,
	}
}

func (c Controller) Create(ctx context.Context, payment Payment) error {
	payment.ID = uuid.NewString()
	return c.kafka.WritePayment(ctx, payment)
}

func (c Controller) GetByID(id string) (Payment, error) {
	return Payment{}, nil
}

func (c Controller) Get() ([]Payment, error) {
	return nil, nil
}

func (c Controller) Update(id string) error {
	return nil
}

func (c Controller) Delete(id string) error {
	return nil
}
