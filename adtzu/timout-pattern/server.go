package main

import (
	"fmt"
	"net/http"
	"time"
)

func simulateHandler(w http.ResponseWriter, r *http.Request) {
	select {
	case <-time.After(5 * time.Second): // Simulate 5 seconds of work
		fmt.Fprintln(w, "Process completed successfully!")
	case <-r.Context().Done(): // Handle client timeout or cancellation
		http.Error(w, "Request canceled or timed out", http.StatusRequestTimeout)
	}
}

func main() {
	http.HandleFunc("/simulate", simulateHandler)

	server := &http.Server{
		Addr:              ":8080",
		ReadHeaderTimeout: 2 * time.Second,
	}

	fmt.Println("Server started at :8080")
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		fmt.Printf("Error starting server: %v\n", err)
	}
}
