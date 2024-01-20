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

func useHandler(fn func(w http.ResponseWriter, r *http.Request) error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := fn(w, r); err != nil {
			handleApiErr(w, err)
		}
	}
}

func handleApiErr(w http.ResponseWriter, err error) error {
	slog.Error(err.Error())
	apiErr := &controller.ApiError{
		Status:  http.StatusInternalServerError,
		Message: "internal server error",
	}
	if errors.As(err, &apiErr) {
		return res.JSON(w, apiErr.Status, apiErr)
	}
	validationErr := &schemas.ValidationError{}
	if errors.As(err, &validationErr) {
		return res.JSON(w, http.StatusUnprocessableEntity, &validationErrorResponse{
			Errors: validationErr.Errors,
			Status: http.StatusUnprocessableEntity,
		})
	}
	if errors.Is(err, service.ErrNotFound) {
		apiErr.Status = http.StatusNotFound
		apiErr.Message = err.Error()
	}
	return res.JSON(w, apiErr.Status, apiErr)
}

type validationErrorResponse struct {
	Errors map[string][]string `json:"errors"`
	Status int                 `json:"status"`
}
