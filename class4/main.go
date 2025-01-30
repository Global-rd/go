package main

import "fmt"

type UserCreateFn func(int, string) (int, string)

func main() {
	fruits := []string{"apple", "banana", "pineapple"}
	ActivateUser(fruits...)
}

func ActivateUser(filters ...string) {
	for _, filter := range filters {
		fmt.Print(filter + " ")
	}
}

func GetCreateUser() UserCreateFn {

	a := 3

	return func(i int, s string) (int, string) {
		return i + a, s
	}
}

func GetCounter() func() int {
	counter := 0

	return func() int {
		counter++
		return counter
	}
}

func GetPtr(value int) *int {
	return &value
}

func CreateUser(apple int, banana string) (int, string) {

	return apple, banana
}
