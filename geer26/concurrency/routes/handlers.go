package routes

import (
	"fmt"
	"net/http"
)

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello world")
}

func IntFetcher(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, 123)
}

func PrimeFetcher(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, 3)
}

func TimeoutFetcher(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Error")
}
