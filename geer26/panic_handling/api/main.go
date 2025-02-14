package main

import (
	"fmt"
	"math/rand"
	"net/http"
)

const (
	RANDCHARS = "0123456789A"
)

func random_selection() string {
	randomIndex := rand.Intn(len(RANDCHARS))
	return string(RANDCHARS[randomIndex])
}

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, fmt.Sprintf("%s", random_selection()))
	})

	http.ListenAndServe(":5000", nil)
}
