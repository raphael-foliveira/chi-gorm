package routes

import (
	"errors"
	"log/slog"
	"net/http"

	"github.com/raphael-foliveira/chi-gorm/internal/exceptions"
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
	apiErr := &exceptions.ApiError{}
	if errors.As(err, &apiErr) {
		res.JSON(w, apiErr.Status, apiErr)
		return
	}
	if errors.Is(err, service.ErrNotFound) {
		err := exceptions.NotFound(err.Error())
		res.JSON(w, err.Status, err)
		return
	}
	res.JSON(w, http.StatusInternalServerError, exceptions.InternalServerError("internal server error"))
}
