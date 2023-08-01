package resp

import (
	"encoding/json"
	"net/http"
)

type M map[string]interface{}

func JSON(w http.ResponseWriter, status int, data interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(data)
}
