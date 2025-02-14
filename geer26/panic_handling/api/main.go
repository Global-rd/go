package main

import (
	"fmt"
	"math/rand"
	"net/http"
)

const (
	RANDCHARS     = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvxxyz"
	RANDOM_LENGHT = 10
)

func random_selection() string {
	buff := []byte{}
	i := 0
	for i < RANDOM_LENGHT {
		randomIndex := rand.Intn(len(RANDCHARS))
		buff = append(buff, []byte(RANDCHARS)[randomIndex])
		i++
	}
	return string(buff)
}

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, fmt.Sprintf("%s", random_selection()))
	})

	http.ListenAndServe(":5000", nil)
}
