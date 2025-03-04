package main

type functions interface {
}

type User struct {
	functions

	fn   func()
	Name string
}
