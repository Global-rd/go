package routes

import (
	"main/middlewares"
	"net/http"
)

func Attachroutes(m *http.ServeMux) {

	m.HandleFunc("/hello_world", middlewares.AttachMiddlewares(HelloWorld))
	m.HandleFunc("/books", middlewares.AttachMiddlewares(HandleBooks))
	m.HandleFunc("/books/{id}", middlewares.AttachMiddlewares(HandleBooks))

}
