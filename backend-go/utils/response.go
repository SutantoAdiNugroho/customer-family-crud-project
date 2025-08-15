package utils

import (
	"encoding/json"
	"net/http"
)

type ResponseData struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Errors  interface{} `json:"errors,omitempty"`
}

func SuccessResponse(w http.ResponseWriter, status int, message string, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	response := ResponseData{
		Status:  status,
		Message: message,
		Data:    data,
	}

	json.NewEncoder(w).Encode(response)
}

func ErrorResponse(w http.ResponseWriter, status int, message string, errors interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	response := ResponseData{
		Status:  status,
		Message: message,
		Errors:  errors,
	}

	json.NewEncoder(w).Encode(response)
}
