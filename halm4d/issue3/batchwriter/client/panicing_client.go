package client

import (
	"io"
	"net/http"
)

type PanickingClient struct {
	url string
}

func NewPanickingClient(serverUrl string) *PanickingClient {
	return &PanickingClient{
		url: serverUrl,
	}
}

func (c *PanickingClient) Get() (bodyString string, err error) {
	resp, err := http.Get(c.url)
	if err != nil {
		return bodyString, err
	}
	defer func() {
		err = resp.Body.Close()
	}()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return bodyString, err
	}
	return string(body), err
}
