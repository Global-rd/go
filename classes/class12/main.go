package main

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	handleJsonFunc := http.HandlerFunc(handleJSON)
	handlePostFunc := http.HandlerFunc(handlePost)

	mux.HandleFunc("/", handleGetRoot)
	mux.HandleFunc("/query", handleGetQuery)
	mux.HandleFunc("/user/{userID}", loggingMiddleware(handlePostFunc))
	mux.HandleFunc("/json", loggingMiddleware(handleJsonFunc))

	slog.Info("server started at 8080")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		slog.Error(err.Error())
	}
}

func loggingMiddleware(next http.Handler) func(w http.ResponseWriter, r *http.Request) {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		slog.Info("request", "method", r.Method, "url", r.URL.Path)

		next.ServeHTTP(w, r)
	})
}

func handleGetRoot(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	slog.Info(r.Method)
	fmt.Fprint(w, r.Method)
	fmt.Fprint(w, "hello")
}

func handlePost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	userID := r.PathValue("userID")
	slog.Info(userID)

	fmt.Fprint(w, userID)
}

type User struct {
	Name string `json:"first_name"`
	Age  int    `json:"age"`
}

func handleJSON(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid function body", http.StatusMethodNotAllowed)
		slog.Error(err.Error())
		return
	}

	defer r.Body.Close()

	user.Name = fmt.Sprintf("Modified: %s", user.Name)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)

	return
}

func handleGetQuery(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		slog.Error("Invalid request method", "code", http.StatusMethodNotAllowed)
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	filterValue := r.URL.Query().Get("filter")
	slog.Info("received message", "filter-value", filterValue)

	response := map[string]string{"message": fmt.Sprintf("Received filter: %s", filterValue)}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
