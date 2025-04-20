package main

import (
	"async-caller/async"
	"async-caller/timeout"
	"errors"
	"log/slog"
	"strconv"
	"time"
)

func init() {
	slog.SetLogLoggerLevel(slog.LevelDebug)
}

func main() {

	client := async.NewRestClient(
		"delayer",
		"http://localhost:8080/delay/{delay}",
		100*time.Second, // Set a longer timeout to test the timeout group
	)

	timeoutGroup := timeout.NewTimeoutGroup[async.RestApiResponseTask](5 * time.Second)
	for i := 0; i < 10; i++ {
		job := func() (async.RestApiResponseTask, error) {
			return client.SimpleCallRestApi(strconv.Itoa(i))
		}
		timeoutGroup.AddJob(timeout.AsyncJob[async.RestApiResponseTask]{
			JobId: strconv.Itoa(i),
			Job:   job,
		})
	}

	responses := timeoutGroup.WaitAll()

	for response := range responses {
		if errors.Is(response.Error, timeout.ErrTimeout) {
			slog.Warn(
				"Timeout calling rest api",
				"jobId", response.JobId,
				"error", response.Error,
			)
			continue
		} else if response.Error != nil {
			slog.Error(
				"Error calling rest api",
				"jobId", response.JobId,
				"error", response.Error,
			)
			continue
		}
		slog.Info(
			"Rest api call completed",
			"jobId", response.JobId,
			"taskId", response.Result.TaskId,
			"url", response.Result.Url,
			"response", response.Result.Response,
			"error", response.Error,
		)
	}

}
