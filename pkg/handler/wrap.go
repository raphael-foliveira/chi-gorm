package handler

import (
	"encoding/json"
	"net/http"
)

func Wrap(fn func(w http.ResponseWriter, r *http.Request) error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := fn(w, r)
		if err != nil {
			json.NewEncoder(w).Encode(map[string]interface{}{
				"error":  err.Error(),
				"status": w.Header().Get("status"),
			})
			return
		}
	}
}
