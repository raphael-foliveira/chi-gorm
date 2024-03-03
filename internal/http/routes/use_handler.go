package routes

import (
	"errors"
	"log/slog"
	"net/http"

	"github.com/raphael-foliveira/chi-gorm/internal/exceptions"
	"github.com/raphael-foliveira/chi-gorm/internal/http/controller"
	"github.com/raphael-foliveira/chi-gorm/internal/http/res"
	"github.com/raphael-foliveira/chi-gorm/internal/validate"
)

func useHandler(fn controller.ControllerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := fn(w, r)
		if err != nil {
			handleApiErr(w, err)
		}
	}
}

func handleApiErr(w http.ResponseWriter, err error) error {
	slog.Error(err.Error())
	apiErr := &exceptions.ApiError{
		Status:  http.StatusInternalServerError,
		Message: "internal server error",
	}
	validationErr := &validate.ValidationError{}
	if errors.As(err, &validationErr) {
		return res.JSON(w, http.StatusUnprocessableEntity, validationErr)
	}
	errors.As(err, &apiErr)
	return res.JSON(w, apiErr.Status, apiErr)
}
