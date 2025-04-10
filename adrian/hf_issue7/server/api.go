package server

import (
	"fmt"
	"github.com/go-chi/chi"
	"log/slog"
	"math/rand"
	"net/http"
	"patterns/config"
	"time"
)

// Return a random time in between 1500 - 2500 milliseconds in multiples of 100 ms
func getWaitMilliSeconds() int {
	min := 15
	max := 25
	return (rand.Intn(max-min+1) + min) * 100

}

func NewDummyRouter(logger *slog.Logger) chi.Router {
	router := chi.NewRouter()
	path := fmt.Sprintf("%s{id}", config.SrvPath)
	router.Get(path, func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		logger.Info("processing request", slog.String("id", id))
		time.Sleep(time.Duration(getWaitMilliSeconds()) * time.Millisecond)
		w.WriteHeader(http.StatusOK)
		_, err := w.Write([]byte(fmt.Sprintf("OK: %s", id)))
		if err != nil {
			logger.Error(err.Error())
		}
	})
	return router
}
