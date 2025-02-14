package main

import (
	"client/writer"
	"errors"
	"fmt"
	"io"
	"net/http"
)

const (
	API         = "http://localhost:5000/"
	MAX_RESULTS = 50
)

func ReadApi(w writer.Writer) {
	for range MAX_RESULTS {
		response, err := http.Get(API)

		if err != nil {
			panic(err.Error())
		}

		responseData, err := io.ReadAll(response.Body)
		if err != nil {
			panic(err.Error())
		}

		err = w.Write(string(responseData))
		if err != nil {
			if errors.Is(err, writer.BufferFullError) {
				w.Flush()
			} else if errors.Is(err, writer.ContainsInvalidCharacterError) {
				panic(err.Error())
			} else {
				w.Close()
				panic(errors.New("unexpected error"))
			}

		}
	}

	w.Close()
}

func main() {

	writer := writer.NewWriter()

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered. Error:\n", r)
			writer.Flush()
		}
	}()

	ReadApi(writer)

}
