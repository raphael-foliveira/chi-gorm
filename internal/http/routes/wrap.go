package routes

import (
	"errors"
	"log/slog"
	"net/http"

	"github.com/raphael-foliveira/chi-gorm/internal/http/controller"
	"github.com/raphael-foliveira/chi-gorm/internal/http/res"
	"github.com/raphael-foliveira/chi-gorm/internal/http/schemas"
	"github.com/raphael-foliveira/chi-gorm/internal/service"
)

func wrap(fn func(w http.ResponseWriter, r *http.Request) error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := fn(w, r); err != nil {
			handleApiErr(w, err)
		}
	}
}

func handleApiErr(w http.ResponseWriter, err error) {
	slog.Error(err.Error())
	apiErr := &controller.ApiError{}
	errValidation := &schemas.ValidationError{}
	if errors.As(err, &apiErr) {
		res.JSON(w, apiErr.Status, apiErr)
		return
	}
	if errors.As(err, &errValidation) {
		res.JSON(w, errValidation.Status, errValidation)
		return
	}
	if errors.Is(err, service.ErrNotFound) {
		err := controller.NotFound(err.Error())
		res.JSON(w, err.Status, err)
		return
	}
	res.JSON(w, http.StatusInternalServerError, controller.InternalServerError("internal server error"))
}
