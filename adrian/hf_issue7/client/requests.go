package client

import (
	"fmt"
	"github.com/google/uuid"
	"io"
	"log/slog"
	"net/http"
	"patterns/config"
	"patterns/model"
	"time"
)

type Request struct {
	logger  *slog.Logger
	baseUrl string
}

func NewRequest(logger *slog.Logger, cfg *config.ServerCfg) *Request {
	return &Request{
		logger:  logger,
		baseUrl: fmt.Sprintf("http://%s:%d%s", cfg.Address, cfg.Port, config.SrvPath),
	}
}

func (r *Request) callResource(resultChan chan<- model.Result, resourceId string) {
	defer close(resultChan)
	resourceUrl := fmt.Sprintf("%s%s", r.baseUrl, resourceId)
	r.logger.Info("calling resource", slog.String("url", resourceUrl))
	resp, err := http.Get(resourceUrl)
	if err != nil {
		r.logger.Error(err.Error())
		resultChan <- model.Result{
			ResourceId: resourceId,
			Status:     ResultNok,
		}
		return
	}
	if resp.StatusCode != http.StatusOK {
		r.logger.Error(fmt.Sprintf("resource call failed with status code: %d", resp.StatusCode))
		resultChan <- model.Result{
			ResourceId: resourceId,
			Status:     ResultNok,
		}
		return
	}
	defer resp.Body.Close()

	var body []byte
	body, err = io.ReadAll(resp.Body)
	if err != nil {
		resultChan <- model.Result{
			ResourceId: resourceId,
			Status:     ResultNok,
		}
		return
	}
	r.logger.Info(fmt.Sprintf("resource called: %s with repsonse: %s", resourceId, string(body)))
	resultChan <- model.Result{
		ResourceId: resourceId,
		Status:     ResultOk,
	}
}

func (r *Request) CallResourceWithTimeout(resultChan chan<- model.Result) {
	defer close(resultChan)
	tmpChan := make(chan model.Result)
	resourceId := uuid.NewString()
	go r.callResource(tmpChan, resourceId)
	select {
	case result := <-tmpChan:
		resultChan <- result
	case <-time.After(time.Second * RequestTimeoutSeconds):
		r.logger.Error(fmt.Sprintf("resource call timed out for id: %s", resourceId))
		resultChan <- model.Result{
			ResourceId: resourceId,
			Status:     ResultTimeout,
		}
	}
}
