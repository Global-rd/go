package routes

import (
	"fmt"
	"net/http"
)

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello world")
}
