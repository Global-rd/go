package main

import (
	"client/writer"
	"errors"
	"fmt"
	"io"
	"net/http"
)

const (
	// Hardcoded API endpoint for fet randon strings
	API = "http://localhost:5000/"
)

func ReadApi(w writer.Writer) {
	// Reads API, and registers strings endlessly
	for {
		response, err := http.Get(API)

		if err != nil {
			panic(err.Error())
		}

		responseData, err := io.ReadAll(response.Body)
		if err != nil {
			panic(err.Error())
		}

		// Buffers incoming strings
		err = w.Write(string(responseData))
		if err != nil {
			if errors.Is(err, writer.BufferFullError) {
				// Panic event when buffer is full (should be handled differently!)
				panic(err.Error())
			} else if errors.Is(err, writer.ContainsInvalidCharacterError) {
				// Panics when the input is corrupted (contains an invalid xharaxter)
				panic(err.Error())
			} else {
				// Panics also at any other unexpepcted error,
				// besides closes the writer
				w.Close()
				panic(errors.New("unexpected error"))
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

	ReadApi(writer)

}
