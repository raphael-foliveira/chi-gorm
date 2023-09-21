package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/raphael-foliveira/chi-gorm/pkg/http/res"
)

func getIdFromPath(r *http.Request) (int64, error) {
	id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		fmt.Println(err)
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
		fmt.Println(err)
		return res.ApiError{
			Message: "invalid body",
			Status:  http.StatusBadRequest,
		}
	}
	return nil
}
