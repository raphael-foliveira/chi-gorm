package res

import (
	"encoding/json"
	"net/http"
)

type response struct {
	w http.ResponseWriter
	s int
}

func New(w http.ResponseWriter) *response {
	return &response{w, http.StatusOK}
}

func (r *response) Status(s int) *response {
	r.s = s
	r.w.WriteHeader(r.s)
	return r
}

func (r *response) JSON(data interface{}) error {
	return json.NewEncoder(r.w).Encode(data)
}

func (r *response) Error(message string) error {
	return r.JSON(ApiError{
		Message: message,
		Status:  r.s,
	})
}

func (r *response) Send() error {
	return nil
}
