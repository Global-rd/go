package routes

import (
	"net/http"
)

func Helloroute(w http.ResponseWriter, r *http.Request) {
	err := DbConnection.Ping()
	if err != nil {
		panic("DATABASE PING ERROR IN HELLOROUTE!")
	}
	w.Write([]byte("Hello World!"))
}
