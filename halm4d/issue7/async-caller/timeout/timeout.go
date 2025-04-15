package timeout

import (
	"context"
	"errors"
	"log/slog"
	"sync"
	"time"
)

type Group[T any] struct {
	Context context.Context
	Timeout time.Duration
	Jobs    []AsyncJob[T]
}

func NewTimeoutGroup[T any](ctx context.Context, timeout time.Duration) *Group[T] {
	return &Group[T]{
		Context: ctx,
		Timeout: timeout,
		Jobs:    make([]AsyncJob[T], 0),
	}
}

type AsyncJob[T any] struct {
	JobId string
	Job   func() (T, error)
}

type AsyncJobResult[T any] struct {
	JobId  string
	Result *T
	Error  error
}

func (st *Group[T]) AddJob(job AsyncJob[T]) {
	st.Jobs = append(st.Jobs, job)
}

func (st *Group[T]) WaitAll() <-chan AsyncJobResult[T] {
	var wg sync.WaitGroup
	wg.Add(len(st.Jobs))

	results := make(chan AsyncJobResult[T], len(st.Jobs))
	for _, job := range st.Jobs {
		go func() {
			ctx, cancel := context.WithTimeout(st.Context, st.Timeout)
			defer cancel()
			defer wg.Done()

			slog.Debug("Executing job", "jobId", job.JobId)
			resultChannel := make(chan AsyncJobResult[T], 1)

			go func() {
				result, err := job.Job()
				if err != nil {
					resultChannel <- AsyncJobResult[T]{
						JobId:  job.JobId,
						Result: &result,
						Error:  err,
					}
				} else {
					resultChannel <- AsyncJobResult[T]{
						JobId:  job.JobId,
						Result: &result,
						Error:  err,
					}
				}
			}()

			select {
			case result := <-resultChannel:
				results <- result
			case <-ctx.Done():
				results <- AsyncJobResult[T]{
					JobId:  job.JobId,
					Result: nil,
					Error:  errors.New("timeout"),
				}
				return
			case <-st.Context.Done():
				results <- AsyncJobResult[T]{
					JobId:  job.JobId,
					Result: nil,
					Error:  st.Context.Err(),
				}
				return
			}
		}()
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	return results
}
