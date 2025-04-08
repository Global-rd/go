package main

import (
	"log"
	"main/routes"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	downstream := http.NewServeMux()
	routes.Attachmhain(downstream)
	api1 := http.NewServeMux()
	routes.AttachIntFetcher(api1)
	api2 := http.NewServeMux()
	routes.AttachPrimeFetcher(api2)
	api_timeout := http.NewServeMux()
	routes.AttachTimeoutFetcher(api_timeout)

	go func() {
		log.Println("Main endpoint started on port 5000...")
		err := http.ListenAndServe(":5000", downstream)
		if err != nil {
			log.Fatal("main endpoint start error")
		}
	}()

	go func() {
		log.Println("Int generator endpoint started on port 5001...")
		err := http.ListenAndServe(":5001", api1)
		if err != nil {
			log.Fatal("int generator endpoint start error")
		}
	}()

	go func() {
		log.Println("Prime generator endpoint started on port 5002...")
		err := http.ListenAndServe(":5002", api2)
		if err != nil {
			log.Fatal("prime generator endpoint start error")
		}
	}()

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
