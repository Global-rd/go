package main

import (
	"context"
	"fmt"
	"time"
)

func resourceCall(ctx context.Context, resourceName string) chan string {
	result := make(chan string)

	go func() {
		defer close(result)
		select {
		case <-ctx.Done():
			return
		case <-time.After(1 * time.Second):
			result <- fmt.Sprintf("%s completed", resourceName)
		}
	}()

	return result
}

func main() {
	resources := []string{"Resource1", "Resource2", "Resource3"}
	timeout := 3 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	results := make(chan string)
	for _, resource := range resources {
		go func(resource string) {
			for res := range resourceCall(ctx, resource) {
				results <- res
			}
		}(resource)
	}

	for i := 0; i < len(resources); i++ {
		select {
		case result := <-results:
			fmt.Println(result)
		case <-ctx.Done():
			fmt.Println("Timeout occurred")
			return
		}
	}

	fmt.Println("All resource calls completed")
}
