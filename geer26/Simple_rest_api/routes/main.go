package routes

import (
	"main/middlewares"
	"net/http"
)

func Attachroutes(m *http.ServeMux) error {

	m.HandleFunc("/hello_world", middlewares.AttachMiddlewares(HelloWorld))

	//handleJsonFunc := http.HandlerFunc(handleJSON)
	//handlePostFunc := http.HandlerFunc(handlePost)
	//mux.HandleFunc("/", handleGetRoot)
	//mux.HandleFunc("/query", handleGetQuery)
	//mux.HandleFunc("/user/{userID}", loggingMiddleware(handlePostFunc))
	//mux.HandleFunc("/json", loggingMiddleware(handleJsonFunc))

	return nil
}
