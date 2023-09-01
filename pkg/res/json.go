package res

import (
	"encoding/json"
	"net/http"
)

type M map[string]interface{}

func JSON(w http.ResponseWriter, status int, data interface{}) error {
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(data)
}
