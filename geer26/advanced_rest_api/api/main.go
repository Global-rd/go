package main

import (
	"advrest/routes"
	"log/slog"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	routes.Attachroutes(mux)

	/*
		if _, err := database.DialStore(); err != nil {
			slog.Error(err.Error())
			panic("Quitting server...")
		}
	*/

	slog.Info("Server started at 5000")
	err := http.ListenAndServe(":5000", mux)
	if err != nil {
		slog.Error(err.Error())
	}
}
