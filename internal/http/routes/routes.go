package routes

import (
	"errors"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/raphael-foliveira/chi-gorm/internal/exceptions"
	"github.com/raphael-foliveira/chi-gorm/internal/http/controller"
	"github.com/raphael-foliveira/chi-gorm/internal/validate"
)

type router struct {
	*chi.Mux
}

func (r *router) Get(path string, fn ControllerFunc) {
	r.Mux.Get(path, useHandler(fn))
}

func (r *router) Post(path string, fn ControllerFunc) {
	r.Mux.Post(path, useHandler(fn))
}

func (r *router) Put(path string, fn ControllerFunc) {
	r.Mux.Put(path, useHandler(fn))
}

func (r *router) Patch(path string, fn ControllerFunc) {
	r.Mux.Patch(path, useHandler(fn))
}

func (r *router) Delete(path string, fn ControllerFunc) {
	r.Mux.Delete(path, useHandler(fn))
}

type ControllerFunc func(*controller.Context) error

func useHandler(fn ControllerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		context := controller.NewContext(w, r)
		err := fn(context)
		if err != nil {
			handleApiErr(context, err)
		}
	}
}

func handleApiErr(ctx *controller.Context, err error) error {
	slog.Error(err.Error())
	apiErr := &exceptions.ApiError{
		Status:  http.StatusInternalServerError,
		Message: "internal server error",
	}
	validationErr := &validate.ValidationError{}
	if errors.As(err, &validationErr) {
		return ctx.JSON(http.StatusUnprocessableEntity, validationErr)
	}
	errors.As(err, &apiErr)
	return ctx.JSON(apiErr.Status, apiErr)
}
