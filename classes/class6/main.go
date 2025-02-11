package main

import (
	"fmt"
)

type EventLog struct {
	Big int64

	ID     int8 `json:"id"`
	Name   string
	Age    int8
	Gender bool
}

func (e EventLog) GetID() int8 {
	return e.ID
}

type Book struct {
	EventLog
	Title string
}

func (u Book) GetTitle() string {
	return u.Title
}

type User struct {
	EventLog
	Name string
}

func (u User) GetName() string {
	return u.Name
}

func CreateUser(name string) User {
	return User{
		Name: name,
	}
}

func (u User) GetID() string {
	id := u.GetID()
	return fmt.Sprintf("%s-suffix", id)
}

func embeddingExamples() {
	user := CreateUser("Jakab")

	userID := user.GetID()

	book := Book{
		Title: "test title",
	}

	fmt.Println("userID", userID)
	fmt.Println("name", user.GetName())

	fmt.Println("bookID", book.GetID())
	fmt.Println("title", book.GetTitle())
}

func main() {
	layoutExample()
}

type Apple struct {
	A string
	B *Banana
}

type Banana struct {
	A Apple
	B string
}
