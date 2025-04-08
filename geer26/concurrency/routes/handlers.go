package routes

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"time"
)

type ReturnStruct struct {
	Status  int
	Random  int
	Prime   int
	Timeout string
	Errors  []string
}

type ReturnInt struct {
	Status int
	Result int
}

type ReturnString struct {
	Status int
	Result string
}

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	retval := ReturnStruct{
		Status: 1,
	}
	rch := make(chan int)
	pch := make(chan int)
	sch := make(chan string)
	//var errs []error

	go func() {
		defer close(rch)
		resp, err := http.Get("http://localhost:5001/")
		if err != nil {
			retval.Status = -1
			retval.Random = -1
			retval.Errors = append(retval.Errors, err.Error())
			return
		}
		var randint ReturnInt
		err = json.NewDecoder(resp.Body).Decode(&randint)
		retval.Random = randint.Result
	}()

	go func() {
		defer close(pch)
		resp, err := http.Get("http://localhost:5002/")
		if err != nil {
			retval.Status = -1
			retval.Prime = -1
			retval.Errors = append(retval.Errors, err.Error())
			return
		}
		var prime ReturnInt
		err = json.NewDecoder(resp.Body).Decode(&prime)
		retval.Prime = prime.Result
	}()

	go func() {
		defer close(sch)
		resp, err := http.Get("http://localhost:5003/")
		if err != nil {
			retval.Status = -1
			retval.Timeout = ""
			retval.Errors = append(retval.Errors, err.Error())
			return
		}
		var str ReturnString
		err = json.NewDecoder(resp.Body).Decode(&str)
		retval.Timeout = str.Result
	}()

	go func() {
		select {
		case <-time.After(time.Millisecond * 2000):
			retval.Status = -1
			retval.Errors = append(retval.Errors, "timeout error")
			retval.Timeout = ""
			return
		}
	}()

	<-sch
	<-pch
	<-rch

	if retval.Status > 0 {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(retval)
}

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
	var heavyload = rand.Intn(6000)
	time.Sleep(time.Millisecond * time.Duration(heavyload))
	retval := ReturnString{
		Status: 1,
		Result: "Generated string! All OK!",
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(retval)
}
