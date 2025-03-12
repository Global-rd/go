package middlewares

import "net/http"

type Middleware func(http.HandlerFunc) http.HandlerFunc

type MiddlewareStack struct {
	Stack []Middleware
}

func AttachMiddlewares(f http.HandlerFunc) http.HandlerFunc {
	stack := []Middleware{
		Logger(),
	}
	for _, middleware := range stack {
		f = middleware(f)
	}
	return f
}
