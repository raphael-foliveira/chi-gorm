package routes

import (
	"errors"
	"log/slog"
	"net/http"
	"strings"

	"github.com/raphael-foliveira/chi-gorm/internal/http/controller"
	"github.com/raphael-foliveira/chi-gorm/internal/http/res"
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
	apiErr := &controller.ApiError{
		Status:  http.StatusInternalServerError,
		Message: "internal server error",
	}
	errors.As(err, &apiErr)
	if apiErr.Status == http.StatusUnprocessableEntity {
		res.JSON(w, apiErr.Status, map[string]any{
			"errors": strings.Split(apiErr.Message, "\n"),
			"status": apiErr.Status,
		})
		return
	}
	if errors.Is(err, service.ErrNotFound) {
		apiErr.Status = http.StatusNotFound
		apiErr.Message = err.Error()
	}
	res.JSON(w, apiErr.Status, apiErr)
}
