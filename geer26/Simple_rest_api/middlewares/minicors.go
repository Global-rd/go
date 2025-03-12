package middlewares

import (
	"net/http"
)

var MethodWhitelist = []string{"GET", "DELETE", "POST"}
var OriginWhitelist = []string{"http://localhost:5000", "http://localhost:8080"}

func CheckMethod(method string) bool {
	for _, m := range MethodWhitelist {
		if m == method {
			return true
		}
	}
	return false
}

func CheckOrigin(origin string) bool {
	for _, o := range OriginWhitelist {
		if o == origin {
			return true
		}
	}
	return false
}

func MiniCORS() Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			// BEFORE REQUEST

			// Call the next middleware/handler in chain
			f(w, r)
			//AFTER REQUEST, BEFORE DEFERRED
		}
	}
}
