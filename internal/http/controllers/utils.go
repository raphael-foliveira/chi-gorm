package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/raphael-foliveira/chi-gorm/internal/http/res"
)

func getIdFromPath(r *http.Request) (int64, error) {
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		return 0, res.ApiError{
			Message: "invalid id",
			Status:  http.StatusBadRequest,
		}
	}
	return id, nil
}

func parseBody(r *http.Request, v interface{}) error {
	err := json.NewDecoder(r.Body).Decode(v)
	if err != nil {
		return res.ApiError{
			Message: "invalid body",
			Status:  http.StatusBadRequest,
		}
	}
	return nil
}
