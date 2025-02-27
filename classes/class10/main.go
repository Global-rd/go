package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

func main() {
	apiURL := "http://localhost:11434/api/generate"

	reader := strings.NewReader(`{
"model": "llama3.2",
"stream": false,
"prompt": "generate 10 random data in json for books and include the following fields:id, title, author, createdat,color and only return with the json do not add your response"
}`)

	resp, err := http.Post(
		apiURL,
		"application/json",
		reader,
	)
	if err != nil {
		log.Fatalf("Failed to fetch data: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Unexpected status code: %d", resp.StatusCode)
	}

	var response Response
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		log.Fatalf("Failed to parse JSON: %v", err)
	}

	fmt.Println(response.Response)

	fmt.Println("✅ All data returned")
}
