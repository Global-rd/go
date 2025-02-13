package main

import (
	"errors"
	"fmt"
	"os"
	"time"

	errgithub "github.com/pkg/errors"
)

func main() {
	fmt.Println("Learn coding")
	val := panicker()
	fmt.Println("panic", val)

	go func() {
		fmt.Println("go routine")
		panic("error in go routine before while sleep happens on the main")
	}()

	time.Sleep(10 * time.Second)
	fmt.Println("hello after sleep")
}

func unHandledPanic() {
	panic("error")
}

func panicker() string {
	fmt.Println("it will panic")
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Error:", err)
		}
	}()

	text := "hello"
	fmt.Println(text)
	panic("Fatal error")
	fmt.Println("this won't be printed")

	return "hello"
}

func deferFileClose() error {
	file, err := os.Open("/tmp/existing")
	if err != nil {
		return err
	}
	defer file.Close()

	defer func() {
		err := file.Close()
		if err != nil {
			fmt.Println(err)
		}
	}()

	fmt.Println("hello2")

	return file.Close()
}

func divide() {
	// panicMain()
	a := 1
	b := 0

	fmt.Println(a / b)
	fmt.Println("Follow The Pattern!") // won't be executed
}

func panicMain() {
	_, err := os.Open("/tmp/nonexisting")
	if err != nil {
		panic(err) // panic: open /fp/nonexisting: no such file or directory
		fmt.Println("Follow The Pattern!")
	}
}

var ErrNotFound = errors.New("not found")

func errorAs() {
	wrappedErr := fmt.Errorf("operation failed: %w", ErrNotFound)

	if errors.Is(wrappedErr, ErrNotFound) {
		fmt.Println("The error is 'not found'")
	} else {
		fmt.Println("Different error")
	}
}

type CustomError struct {
	Msg string
}

func (e CustomError) Error() string {
	return e.Msg
}

func customError() {

	// var err error = TimeoutError

	err := CustomError{
		Msg: "custom error message",
	}

	wrappedErr := errgithub.Wrap(err, "unable to load user data")

	switch wrappedErr.(type) {
	case CustomError:
		fmt.Println("custom-error")
	case *DBError:
		fmt.Println("db-error")
	default:
		fmt.Println("unknown error")
	}

	// er = &CustomError{"Something went wrong"}
	// var target *CustomError

	if errors.As(wrappedErr, &err) {
		fmt.Println(wrappedErr, "Custom error detected")
	} else {
		fmt.Println("Unknown error from as")
	}

	if errors.Is(wrappedErr, errors.New("unable to load user data")) {
		fmt.Println("error is detected the message itself")
	}

}

func errorF() {
	userID := 42
	err := fmt.Errorf("user with ID %d not found", userID)
	fmt.Println(err)
}

func dbError() {

	conn, err := ConnectDB("postgres://localhost:5432")

	if err != nil {
		fmt.Println("error", err)
		os.Exit(1)
	}

	fmt.Println(conn)
}
