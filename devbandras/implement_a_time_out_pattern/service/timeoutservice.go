package service

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

type TimeOutOperation[T any] struct {
	Ctx            context.Context     // Az operation művelete
	Timeout        time.Duration       // A művelet timeout-ja
	Operation      func() (T, error)   // A végrehajtandó művelet
	ResponseWriter http.ResponseWriter // http response-on végrehajtandó művelet
}

// NewTimeoutOperation létrehoz egy új TimeoutOperation példányt
func NewTimeOutOperation[T any](ctx context.Context, timeout time.Duration, operation func() (T, error), w http.ResponseWriter) *TimeOutOperation[T] {
	return &TimeOutOperation[T]{
		Ctx:            ctx,
		Timeout:        timeout,
		Operation:      operation,
		ResponseWriter: w,
	}
}

// Az Execute időkorláttal hajtja végre a webservice műveletet
// Parameters
//
// Returns:
// - T: a meghívandó függvény típusa
// - bool: a művelet sikeressége
// Ha a művelet időtúllépéssel jár akkor a response timeout választ küld
// Ha a művelet sikertelen akkor a megadott hiba státuszkódot küldi el a response-on
// Ha a művelet sikeres akkor a művelet eredményét adja vissza
func (t *TimeOutOperation[T]) Execute() (T, bool) {
	// Kontextus létrehozása időtúllépéssel
	timeOutCtx, cancel := context.WithTimeout(t.Ctx, t.Timeout)
	defer cancel()

	// Csatorna az eredmény fogadására
	ch := make(chan struct {
		result T
		err    error
	})

	// Goroutine indítása a művelet végrehajtására
	go func() {
		result, err := t.Operation()
		ch <- struct {
			result T
			err    error
		}{result, err}
	}()

	var zero T

	// Fan-in pattern: a goroutine eredményének fogadása
	// Várakozás a goroutine eredményére vagy a timeout-ra
	select {
	// Ha a goroutine befejeződik, akkor az eredményt visszaadjuk
	case result := <-ch:
		if result.err != nil {
			http.Error(t.ResponseWriter, result.err.Error(), http.StatusInternalServerError)
			return zero, false
		}
		// sikeres művelet, visszaadjuk az eredményt
		return result.result, true

	// Ha a művelet időtúllépéssel jár, akkor a response-on küldjük el a timeout választ
	case <-timeOutCtx.Done():
		err := fmt.Errorf("timeout: a művelet nem fejeződött be a megadott időn belül")
		http.Error(t.ResponseWriter, err.Error(), http.StatusGatewayTimeout)
		return zero, false
	}
}
