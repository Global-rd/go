package client

import (
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type PanickingClient struct {
	url       string
	panicRate int
}

func NewPanickingClient(serverUrl string, panicRate int) *PanickingClient {
	return &PanickingClient{
		url:       serverUrl,
		panicRate: panicRate,
	}
}

func (c *PanickingClient) Get() (bodyString string) {
	resp, err := http.Get(c.url)
	if err != nil {
		log.Fatalln(err)
	}
	defer func() {
		err := resp.Body.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	bodyString = string(body)
	requestCount, err := strconv.Atoi(strings.Split(bodyString, ": ")[1])
	if err != nil {
		log.Fatalln(err)
	}
	if requestCount%c.panicRate == 0 {
		panic("Panic!")
	}
	return bodyString
}
