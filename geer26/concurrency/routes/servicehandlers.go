package routes

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"time"
)

func IntFetcher(inbound <-chan int) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		var randint int
		status := 1

		select {
		case v := <-inbound:
			randint = v
		case <-time.After(time.Millisecond * 2000):
			status = 0
			randint = -1
		}

		retval := ReturnInt{
			Status: status,
			Result: randint,
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(retval)
	}
}

func PrimeFetcher(inbound <-chan int) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		var prime int
		status := 1

		select {
		case v := <-inbound:
			prime = v
		case <-time.After(time.Millisecond * 2000):
			status = 0
			prime = -1
		}

		retval := ReturnInt{
			Status: status,
			Result: prime,
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(retval)
	}

}

func TimeoutFetcher(w http.ResponseWriter, r *http.Request) {

	var heavyload = rand.Intn(4000)
	time.Sleep(time.Millisecond * time.Duration(heavyload))
	retval := ReturnString{
		Status: 1,
		Result: "Generated string! All OK!",
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(retval)

}
