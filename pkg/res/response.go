package res

import (
	"encoding/json"
	"net/http"
)

func SendStatus(w http.ResponseWriter, status int) error {
	w.WriteHeader(status)
	return nil
}

func JSON(w http.ResponseWriter, status int, data interface{}) error {
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(data)
}

func Error(w http.ResponseWriter, err error, status int, message string) error {
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(ApiError{
		Message: message,
		Status:  status,
	})
}
