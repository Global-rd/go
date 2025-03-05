package main

import (
	"client/writer"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
)

const (
	// Hardcoded API endpoint for fet randon strings
	API = "http://localhost:5000/"
)

func ReadApi(w writer.Writer) error {
	// Reads API, and registers strings endlessly
	for {
		response, err := http.Get(API)
		if err != nil {
			return err
		}
		responseData, err := io.ReadAll(response.Body)
		if err != nil {
			return err
		}
		// Buffers incoming strings
		err = w.Write(string(responseData))
		if err != nil {
			if errors.Is(err, writer.BufferFullError) {
				return err
			} else if errors.Is(err, writer.ContainsInvalidCharacterError) {
				// Panics when the input is corrupted (contains an invalid character)
				panic(err.Error())
			} else {
				w.Close()
				return errors.New("unexpected error")
			}
		}
	}
}

func main() {
	writer := writer.NewWriter()
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Recovered. Error:\n", err)
			writer.Flush()
		}
	}()
	err := ReadApi(writer)
	if err != nil {
		log.Println("Error occured: ", err)
	} else {
		log.Println("All string read without error!")
	}
}
