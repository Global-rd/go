package routes

import (
	"encoding/json"
	"main/utils"
	"net/http"
	"time"
)

type Signal struct{}

type ReturnStruct struct {
	CatFact        string `json:"cat_fact"`
	UselessFact    string `json:"useless_fact"`
	BullshitExcuse string `json:"bullshit_excuse"`
	Status         int    `json:"status"`
	Errors         []string
}

type ChannelStruct struct {
	Fetcher func() (string, error)
	Message chan string
	Done    bool
	Err     chan error
}

func BaseHandler(w http.ResponseWriter, r *http.Request) {
	var retval = ReturnStruct{
		Status: 1,
	}

	done := make(chan Signal)

	catFact := ChannelStruct{
		Fetcher: utils.GetCatFact,
		Message: make(chan string),
		Done:    false,
	}

	uselessFact := ChannelStruct{
		Fetcher: utils.GetUselessFact,
		Message: make(chan string),
		Done:    false,
	}

	bullshitExcuse := ChannelStruct{
		Fetcher: utils.GetBullshitExcuse,
		Message: make(chan string),
		Done:    false,
	}

	go func() {
		ApplyFetchers(&catFact, &uselessFact, &bullshitExcuse)
	}()

	go func() {
		defer close(done)
		for {
			if catFact.Done && uselessFact.Done && bullshitExcuse.Done {
				return
			}
			select {
			case v := <-catFact.Message:
				retval.CatFact = v
			case v := <-catFact.Err:
				retval.Errors = append(retval.Errors, v.Error())
				retval.Status = -1
			case v := <-uselessFact.Message:
				retval.UselessFact = v
			case v := <-uselessFact.Err:
				retval.Errors = append(retval.Errors, v.Error())
				retval.Status = -1
			case v := <-bullshitExcuse.Message:
				retval.BullshitExcuse = v
			case v := <-bullshitExcuse.Err:
				retval.Errors = append(retval.Errors, v.Error())
				retval.Status = -1
			case <-time.After(time.Millisecond * 2000):
				retval.Status = -1
				retval.Errors = append(retval.Errors, "timed out")
				return
			}
		}
	}()

	<-done

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(retval)
}

func ApplyFetchers(fetchers ...*ChannelStruct) {

	for _, fetcher := range fetchers {
		go func() {
			fetched, err := fetcher.Fetcher()
			if err != nil {
				fetcher.Err <- err
				fetcher.Done = true
				return
			}
			fetcher.Message <- fetched
			fetcher.Done = true
			return
		}()
	}
}
