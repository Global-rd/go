package main

import (
	"context"
	"log"
	"main/routes"
	"main/utils"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	// Create channel for craceful shutdown
	done := make(chan struct{})
	defer close(done)

	// Spin up worker goroutines
	// Generates int endlessly and send to a channel
	RndInt := utils.Producer(context.Background(), done, utils.GenerateInt)
	// Filters random int tha are greate than 300.000.000 and channels to the next stage
	Gt_3M := utils.Filter(context.Background(), done, RndInt, utils.FilterFunc)
	// Filters out primes and redirects to a channel to consume
	Primes := utils.Filter(context.Background(), done, Gt_3M, utils.CheckifPrime)

	// Prepare API's
	downstream := http.NewServeMux()
	routes.Attachmhain(downstream)
	api1 := http.NewServeMux()
	routes.AttachIntFetcher(api1, RndInt)
	api2 := http.NewServeMux()
	routes.AttachPrimeFetcher(api2, Primes)
	api_timeout := http.NewServeMux()
	routes.AttachTimeoutFetcher(api_timeout)

	//Spin up API's
	//Main API for user interaction
	go func() {
		log.Println("Main endpoint started on port 5000...")
		err := http.ListenAndServe(":5000", downstream)
		if err != nil {
			log.Fatal("main endpoint start error")
		}
	}()

	//API that always return a random int
	go func() {
		log.Println("Int generator endpoint started on port 5001...")
		err := http.ListenAndServe(":5001", api1)
		if err != nil {
			log.Fatal("int generator endpoint start error")
		}
	}()

	// API that returns a prime between 300.000.000 - 500.000.000
	go func() {
		log.Println("Prime generator endpoint started on port 5002...")
		err := http.ListenAndServe(":5002", api2)
		if err != nil {
			log.Fatal("prime generator endpoint start error")
		}
	}()

	// API that tries to return within 2 seconds (using random time.sleep to simulate heavy lifting)
	go func() {
		log.Println("Timeout-ish endpoint started on port 5003...")
		err := http.ListenAndServe(":5003", api_timeout)
		if err != nil {
			log.Fatal("timeout-ish endpoint start error")
		}
	}()

	shutdownChan := make(chan os.Signal, 1)
	signal.Notify(shutdownChan, syscall.SIGINT, syscall.SIGTERM)
	<-shutdownChan

}
