package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type Signal int

type Number interface {
	int64 | float64 | int
}

func Create() {
	var apple DB[int]

	fmt.Println(apple)
}

type DB[T Number] interface {
	GetValue() T
}

type SQLDB[T constraints.Integer] struct {
	value T
}

func (s SQLDB[T]) GetValue() T {
	return s.value + s.value
}
