package httpresponse

import (
	"encoding/json"
	"net/http"
)

func WriteResponseBody(w http.ResponseWriter, statusCode int, body any) {
	w.WriteHeader(statusCode)
	encodeErr := json.NewEncoder(w).Encode(body)
	if encodeErr != nil {
		http.Error(w, "Error encoding error response", http.StatusInternalServerError)
	}
}
