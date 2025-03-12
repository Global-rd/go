package main

import (
	"fmt"
	"log/slog"
	"main/routes"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	if err := routes.Attachroutes(mux); err != nil {
		slog.Error(fmt.Sprintf("error at attaching routes: %s", err.Error()))
		panic("Quitting...")
	}

	slog.Info("server started at 5000")
	err := http.ListenAndServe(":5000", mux)
	if err != nil {
		slog.Error(err.Error())
	}
}
