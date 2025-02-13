package main

import "fmt"

var (
	TimeoutError = &DBError{Code: 43, Message: "connection timed out"}
	// TimeoutError = &DBError{Code: 43, Message: "connection timed out"}
)

type DBConnection struct{}

func ConnectDB(addr string) (DBConnection, error) {
	// connecting

	return DBConnection{}, TimeoutError
}

type DBError struct {
	Code    int
	Message string
}

func (e DBError) Error() string {
	return fmt.Sprintf("code: %d, message: %s", e.Code, e.Message)
}
