package async

import (
	"errors"
	"io"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type RestApiResponseTask struct {
	TaskId   string
	Url      string
	Response string
}

type RestClient struct {
	Name   string
	Url    string
	Client *http.Client
}

func NewRestClient(name string, url string, timeout time.Duration) *RestClient {
	return &RestClient{
		Name: name,
		Url:  url,
		Client: &http.Client{
			Timeout: timeout,
		},
	}
}

func (c *RestClient) SimpleCallRestApi(taskId string) (RestApiResponseTask, error) {
	delay := strconv.Itoa(rand.Intn(10))
	url := strings.Replace(c.Url, "{delay}", delay, 1)
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return RestApiResponseTask{
			TaskId: taskId,
			Url:    url,
		}, err
	}

	resp, err := c.Client.Do(req)
	if err != nil {
		return RestApiResponseTask{
			TaskId: taskId,
			Url:    url,
		}, err
	}

	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return RestApiResponseTask{
			TaskId:   taskId,
			Url:      url,
			Response: string(body),
		}, nil
	} else {
		return RestApiResponseTask{
			TaskId: taskId,
			Url:    url,
		}, errors.New("non-200 response")
	}
}
