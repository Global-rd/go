package main

import (
	"advrest/service"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	service := service.ServiceBuilder().
		Configure().
		Connect().
		AttachRoutes()

	defer service.Db.Close()

	go func() {
		_, err := service.Run()
		if err != nil {
			log.Fatalf("error at service startup: %s", service.InitError.Error())
		}
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan
	log.Println("service stopped")

}
