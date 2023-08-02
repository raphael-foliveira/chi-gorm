package res

import (
	"encoding/json"
	"net/http"
)

func Error(w http.ResponseWriter, status int, message string) error {
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(map[string]interface{}{
		"error": message,
	})
}
