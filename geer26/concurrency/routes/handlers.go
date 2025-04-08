package routes

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type MainStruct struct {
	status  int
	random  int
	prime   int
	timeout string
}

type ReturnInt struct {
	status int
	result int
}

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	// var randint int
	//var prime int
	//var timeout string
	go func() (*http.Response, error) {
		resp, err := http.Get("http://localhost:5001/")
		if err != nil {
			return nil, err
		}
		log.Println(resp)
		return resp, nil
	}()
	fmt.Fprintln(w, "hello world")
}

func IntFetcher(inbound <-chan int) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		random := <-inbound
		retval := ReturnInt{
			status: 1,
			result: random,
		}
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(retval)
	}
}

func PrimeFetcher(inbound <-chan int) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		prime := <-inbound
		fmt.Fprintln(w, prime)
	}
}

func TimeoutFetcher(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Error")
}
