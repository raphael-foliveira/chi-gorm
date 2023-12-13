package res

import (
	"encoding/json"
	"net/http"

	"github.com/raphael-foliveira/chi-gorm/internal/exceptions"
)

func SendStatus(w http.ResponseWriter, status int) error {
	w.WriteHeader(status)
	return nil
}

func JSON(w http.ResponseWriter, status int, data interface{}) error {
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(data)
}

func Error(w http.ResponseWriter, status int, message string) error {
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(exceptions.ApiError{
		Message: message,
		Status:  status,
	})
}
