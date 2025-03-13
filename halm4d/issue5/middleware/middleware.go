package middleware

import (
	"net/http"
)

type Middleware interface {
	Handle(next http.Handler) http.Handler
}

type Chain struct {
	middlewares []Middleware
}

func NewChain() *Chain {
	return &Chain{
		middlewares: []Middleware{},
	}
}

func (mc *Chain) Handle(next http.Handler) http.Handler {
	for i := range mc.middlewares {
		reversedIndex := len(mc.middlewares) - i - 1
		next = mc.middlewares[reversedIndex].Handle(next)
	}
	return next
}

func (mc *Chain) Use(middleware Middleware) {
	mc.middlewares = append(mc.middlewares, middleware)
}
