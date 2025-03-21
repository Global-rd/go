package httpresponse

import (
	"encoding/json"
	"errors"
	"issue6/db"
	"net/http"
)

type ErrorResponse struct {
	StatusCode int    `json:"status_code"`
	Error      string `json:"error"`
	Details    string `json:"details"`
}

func HandleError(err error) (errorResponse *ErrorResponse, ok bool) {
	if err == nil {
		return nil, true
	}
	switch {
	case errors.Is(err, db.NotFoundError):
		return &ErrorResponse{
			StatusCode: http.StatusNotFound,
			Error:      "Book not found",
			Details:    err.Error(),
		}, false
	default:
		return &ErrorResponse{
			StatusCode: http.StatusInternalServerError,
			Error:      "Internal server error",
			Details:    err.Error(),
		}, false
	}
}

func WriteErrorResponse(w http.ResponseWriter, er ErrorResponse) {
	w.WriteHeader(er.StatusCode)
	encodeErr := json.NewEncoder(w).Encode(er)
	if encodeErr != nil {
		http.Error(w, "Error encoding error response", http.StatusInternalServerError)
	}
}
