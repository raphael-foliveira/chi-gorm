package routes

import (
	"encoding/json"
	"net/http"

	"github.com/raphael-foliveira/chi-gorm/pkg/res"
)

func wrap(fn func(w http.ResponseWriter, r *http.Request) error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := fn(w, r)
		if err != nil {
			json.NewEncoder(w).Encode(res.ApiError{
				Message: err.Error(),
				Status:  http.StatusInternalServerError,
			})
		}
	}
}
