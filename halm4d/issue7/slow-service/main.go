package main

import (
	"log/slog"
	"net/http"
	"strconv"
	"time"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/delay/{delay}", func(w http.ResponseWriter, r *http.Request) {
		slog.Info("Received request", "method", r.Method, "url", r.URL.String())
		delay := r.PathValue("delay")
		if delay == "" {
			http.Error(w, "Delay parameter is required", http.StatusBadRequest)
			return
		}
		delayInt, err := strconv.Atoi(delay)
		if err != nil {
			http.Error(w, "Invalid delay parameter", http.StatusBadRequest)
			return
		}
		slog.Info("Received delay", "delay", delayInt)
		time.Sleep(time.Duration(delayInt) * time.Second)
		w.WriteHeader(http.StatusOK)
		_, err = w.Write([]byte("Delay: " + delay))
		if err != nil {
			http.Error(w, "Error writing response", http.StatusInternalServerError)
			return
		}
	})

	slog.Info("Listening on port 8080")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		panic(err)
	}
}
