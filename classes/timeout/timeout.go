package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

const DefaultTimeOut = time.Second * 3

func send(req string) chan string {
	ch := make(chan string)

	go func() {
		params := url.Values{}
		params.Add("q", req)

		url := fmt.Sprintf("https://followthepattern.net/search?%s", params.Encode())

		resp, err := http.Get(url)
		if err != nil {
			ch <- err.Error()
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			ch <- err.Error()
		}

		time.Sleep(time.Second * 5)

		ch <- string(body)
		close(ch)
	}()

	return ch
}

func GetResource() {
	respCh := send("value")

	select {
	case <-time.After(DefaultTimeOut):
		fmt.Println("timed out")
	case resp := <-respCh:
		fmt.Println("resp", resp)
	}
}
