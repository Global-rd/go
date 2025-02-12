package main

type Entity interface {
	GetName() string
}

type Any interface{}

type GetName func() string

func (fn GetName) GetName() string {
	return fn()
}

type Base struct {
	id string
}

func (b Base) ID() string {
	return b.id
}
