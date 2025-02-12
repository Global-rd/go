package main

import (
	"errors"
	"fmt"
)

func main() {
	// var val any = 1

	value := typeAssertionWithSwitch(1)

	// value := typeAssertion()

	fmt.Println("any-val", value)
}

func Printer(a any) {
	fmt.Println("any-val", a)
}

func typeAssertionWithSwitch(val interface{}) string {
	switch specifiedType := val.(type) {
	case interface {
		GetNothing() string
	}:
		return specifiedType.GetNothing()
	case User:
		return specifiedType.GetName()
	case string:
		return specifiedType
	}
	return "wrong type"
}

func typeAssertion(val any) (string, error) {
	user, ok := val.(User)
	if !ok {
		return "", errors.New("invalid type")
	}

	return user.Name, nil
}

func functionType() {

	var function GetName

	var e Entity = function

	e.GetName()

	user := User{}

	GetID(user)
}

func GetID(entity interface {
	ID() string
}) string {
	return entity.ID()
}

func baisInterfaces() {
	fmt.Println("hello")

	user := User{
		Name: "Peter",
	}

	file := File{
		Name: "Photo1.jpg",
	}

	entities := []Entity{user, file}

	for _, entity := range entities {
		entityPrinter(entity)
	}
}

func NewUser() Entity {
	return User{
		Name: "Peter",
	}
}

func entityPrinter(e Entity) {
	fmt.Println("e", e.GetName())
}
