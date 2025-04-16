package utils

import (
	"context"
	"encoding/json"
	"net/http"
)

type UselessFact struct {
	Id        string `json:"id"`
	Language  string `json:"language"`
	Permalink string `json:"parmalink"`
	Source    string `json:"source"`
	SourceUrl string `json:"source_url"`
	Text      string `json:"text"`
}

func Getuselessfact(ctx context.Context) chan string {
	ch := make(chan string)
	uselessfact := UselessFact{}
	go func() {
		defer close(ch)
		uselessfactUrl := "https://uselessfacts.jsph.pl/api/v2/facts/random?language=en"
		resp, err := http.Get(uselessfactUrl)
		if err != nil {
			close(ch)
		}
		err = json.NewDecoder(resp.Body).Decode(&uselessfact)
		if err != nil {
			close(ch)
		}
		ch <- uselessfact.Text
	}()
	return ch
}
