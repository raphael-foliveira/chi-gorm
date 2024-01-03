package routes

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/raphael-foliveira/chi-gorm/internal/exceptions"
	"github.com/raphael-foliveira/chi-gorm/internal/http/res"
)

func wrap(fn func(w http.ResponseWriter, r *http.Request) error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := fn(w, r); err != nil {
			handleApiErr(w, err)
		}
	}
}

func handleApiErr(w http.ResponseWriter, err error) {
	fmt.Println(err.Error())
	apiErr := &exceptions.ApiError{}
	notFoundErr := &exceptions.NotFoundError{}
	validationErr := &exceptions.ValidationError{}
	if errors.As(err, &apiErr) {
		res.JSON(w, apiErr.Status, apiErr)
		return
	}
	if errors.As(err, &notFoundErr) {
		res.JSON(w, http.StatusNotFound, exceptions.ApiError{
			Message: err.Error(),
			Status:  http.StatusNotFound,
		})
		return
	}
	if errors.As(err, &validationErr) {
		res.JSON(w, http.StatusUnprocessableEntity, exceptions.MultipleApiError{
			Errors: strings.Split(err.Error(), "\n"),
			Status: http.StatusUnprocessableEntity,
		})
		return
	}
	res.JSON(w, http.StatusInternalServerError, exceptions.ApiError{
		Message: "internal server error",
		Status:  http.StatusInternalServerError,
	})
}
