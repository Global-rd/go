package routes

import (
	"context"
	"encoding/json"
	"net/http"
	"sync"
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

func MultiFetch(w http.ResponseWriter, r *http.Request) {
	//start := time.Now()
	done := make(chan struct{})
	ctx := context.Background()

	retval := ReturnStruct{
		Random:  -1,
		Prime:   -1,
		Timeout: "",
		Status:  1,
	}

	apilist := []string{
		"http://localhost:5001",
		"http://localhost:5002",
		"http://localhost:5003",
	}

	go fetchAPI(ctx, &retval, apilist, done)

	select {
	case <-done:
		//log.Println("Early finish! - ", time.Since(start))
	case <-time.After(time.Millisecond * 2000):
		retval.Errors = append(retval.Errors, "timeout error")
		retval.Status = -1
		//log.Println("Timed out! - ", time.Since(start))
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(retval)
}

func fetchAPI(ctx context.Context, retval *ReturnStruct, apilist []string, done chan<- struct{}) {
	wg := sync.WaitGroup{}

	for _, to_fetch := range apilist {
		wg.Add(1)
		go func() {

			defer wg.Done()
			resp, err := http.Get(to_fetch)
			if err != nil {
				retval.Status = -1
				retval.Errors = append(retval.Errors, err.Error())
				return
			}

			if to_fetch == "http://localhost:5003" {
				var retv ReturnString
				err = json.NewDecoder(resp.Body).Decode(&retv)
				if err != nil {
					retval.Errors = append(retval.Errors, err.Error())
					retval.Status = -1
					retval.Timeout = ""
					return
				}
				retval.Timeout = retv.Result
				return
			}

			if to_fetch == "http://localhost:5001" {
				var retv ReturnInt
				err = json.NewDecoder(resp.Body).Decode(&retv)
				if err != nil {
					retval.Errors = append(retval.Errors, err.Error())
					retval.Status = -1
					return
				}
				retval.Random = retv.Result
				return
			}

			if to_fetch == "http://localhost:5002" {
				var retv ReturnInt
				err = json.NewDecoder(resp.Body).Decode(&retv)
				if err != nil {
					retval.Errors = append(retval.Errors, err.Error())
					return
				}
				retval.Prime = retv.Result
				return
			}

		}()
	}

	wg.Wait()

	close(done)

	return
}
