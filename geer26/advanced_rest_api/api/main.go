package main

import (
	"advrest/logger"
	"advrest/service"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	service := service.ServiceBuilder().
		Configure().
		CreateLogger(
			logger.WithLogfile("another_logfile.log"),
		).
		Connect().
		AttachRoutes()

	go func() {
		_, err := service.Run()
		if err != nil {
			log.Fatalf("error at service startup: %s", service.InitError.Error())
		}
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	service.Db.Close()
	service.Logger.INFO("Db connection closed")
	service.Logger.CloseLog()

}
