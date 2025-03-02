package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		fmt.Println("SERVER_PORT environment variable not set, applying default value: 8080")
		port = "8080"
	}

	var requestCount int
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		requestCount++
		_, _ = fmt.Fprint(w, "Request count: ", requestCount)
	})
	fmt.Println("Starting server on :8080")
	err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
