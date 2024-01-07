package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/raphael-foliveira/chi-gorm/internal/exceptions"
)

func getUintPathParam(r *http.Request, paramName string) (uint, error) {
	id, err := strconv.ParseUint(chi.URLParam(r, paramName), 10, 64)
	if err != nil {
		return 0, exceptions.NewBadRequestError(fmt.Sprintf("invalid %s", paramName))
	}
	return uint(id), nil
}

func parseBody[T interface{}](r *http.Request, v *T) (*T, error) {
	err := json.NewDecoder(r.Body).Decode(v)
	if err != nil {
		return nil, exceptions.NewBadRequestError("invalid body")
	}
	return v, nil
}
