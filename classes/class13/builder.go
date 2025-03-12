package main

import "time"

type Car struct {
	Model          string
	Color          string
	ProductionDate time.Time
}

type CarBuilder struct {
	car *Car
}

func NewCarBuilder() *CarBuilder {
	return &CarBuilder{
		car: &Car{},
	}
}

func (cb *CarBuilder) SetModel(model string) *CarBuilder {
	cb.car.Model = model
	return cb
}

func (cb *CarBuilder) SetColor(color string) *CarBuilder {
	cb.car.Color = color
	return cb
}

func (cb *CarBuilder) SetProductionDate(productionDate time.Time) *CarBuilder {
	cb.car.ProductionDate = productionDate
	return cb
}

func (cb CarBuilder) Build() Car {
	return *cb.car
}
