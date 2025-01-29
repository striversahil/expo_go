package utils

import (
	"encoding/json"
	"net/http"
)

// make user &Response{}
type ApiResponse struct {
	Message string `json:"Message"`
	StatusCode  int `json:"StatusCode"`
	Payload    interface{}  `json:"Payload,omitempty"`
}



func NewRespose(w http.ResponseWriter, message string, statusCode int, payload interface{}) {
	ApiResponse := ApiResponse{
		Message: message,
		StatusCode: statusCode,
		Payload: payload,
	}

	jsonResponse, err := json.Marshal(ApiResponse)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
	w.WriteHeader(statusCode)
}

func ErrorResponse(w http.ResponseWriter, error_message string, statusCode int) {
	NewRespose(w, error_message, statusCode , nil)
}