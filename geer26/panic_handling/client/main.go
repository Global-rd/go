package main

import (
	"client/writer"
	"io"
	"net/http"
)

const (
	API = "http://localhost:5000/"
)

func main() {

	writer := writer.NewWriter()

	for {
		response, err := http.Get(API)

		if err != nil {
			panic(err.Error())
		}

		responseData, err := io.ReadAll(response.Body)
		if err != nil {
			panic(err.Error())
		}

		err = writer.Write(string(responseData))
		if err != nil {
			panic(err.Error())
		}
	}

}
