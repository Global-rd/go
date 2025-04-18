package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type ChuckJoke struct {
	Value string `json:"value"`
}

const timeoutDuration = 10 * time.Second

func resourceCall() chan string {
	result := make(chan string)

	go func() {
		var joke ChuckJoke

		url := fmt.Sprintf("https://api.chucknorris.io/jokes/random/")

		res, err := http.Get(url)
		if err != nil {
			result <- fmt.Sprintf("Error fetching resource %v", err)
		}
		//defer res.Body.Close()

		body, err := io.ReadAll(res.Body)
		if err != nil {
			result <- fmt.Sprintf("Error reading response for resource %v", err)
		}

		err = json.Unmarshal(body, &joke)
		if err != nil {
			result <- fmt.Sprintf("unmarshalling json run into an error: %w", err)
		}

		time.Sleep(5 * time.Second) // Simulate a long-running operation

		result <- string(joke.Value)
		close(result)
	}()

	return result
}

func main() {
	response := resourceCall()

	select {
	case <-time.After(timeoutDuration):
		fmt.Println("Timeout occurred")
	case res := <-response:
		fmt.Println("Response received: ", res)
	}
}
