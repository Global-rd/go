package routes

import "net/http"

func Helloroute(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!"))
}
