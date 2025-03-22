package routes

import (
	"log"
	"net/http"
)

func Helloroute(w http.ResponseWriter, r *http.Request) {
	log.Println(Database.Stats())
	err := Database.Ping()
	if err != nil {
		panic("DATABASE PING ERROR IN HELLOROUTE!")
	}
	w.Write([]byte("Hello World!"))
}
