package main

import (
	"advrest/service"
	"log"
)

func main() {

	servie, err := service.ServiceBuilder().Configure().Connect().AttachRoutes().Run()
	if err != nil {
		log.Fatalf("error at service startup: %s", err.Error())
	}
	defer servie.Db.Close()

}
