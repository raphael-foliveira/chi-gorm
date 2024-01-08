package routes

import (
	"errors"
	"log/slog"
	"net/http"

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
	slog.Error(err.Error())
	apiErr := &exceptions.ApiError{}
	if errors.As(err, &apiErr) {
		res.JSON(w, apiErr.Status, apiErr)
		return
	}
	res.JSON(w, http.StatusInternalServerError, exceptions.InternalServerError("internal server error"))
}
