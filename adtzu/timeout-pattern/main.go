package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"
)

func fetchWithTimeout(url string, timeout time.Duration) (string, error) {

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return "", fmt.Errorf("request creation failed: %w", err)
	}

	// Execute HTTP request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		// Check if error is due to timeout
		if ctx.Err() == context.DeadlineExceeded {
			return "", fmt.Errorf("request timed out after %v", timeout)
		}
		return "", fmt.Errorf("request failed: %w", err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response: %w", err)
	}

	return string(body), nil
}

func main() {

	url := "https://httpbin.org/delay/3" // Simulates 3-second delay
	timeout := 2 * time.Second

	result, err := fetchWithTimeout(url, timeout)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Response received:\n%s\n", result)
}
