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
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(data)
}
