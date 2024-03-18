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

type Router struct {
	*chi.Mux
}

func (r *Router) Get(path string, fn controller.ControllerFunc) {
	r.Mux.Get(path, useHandler(fn))
}

func (r *Router) Post(path string, fn controller.ControllerFunc) {
	r.Mux.Post(path, useHandler(fn))
}

func (r *Router) Put(path string, fn controller.ControllerFunc) {
	r.Mux.Put(path, useHandler(fn))
}

func (r *Router) Patch(path string, fn controller.ControllerFunc) {
	r.Mux.Patch(path, useHandler(fn))
}

func (r *Router) Delete(path string, fn controller.ControllerFunc) {
	r.Mux.Delete(path, useHandler(fn))
}

func useHandler(fn controller.ControllerFunc) http.HandlerFunc {
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

type Routes struct {
	ClientsRoutes     *Router
	OrdersRoutes      *Router
	ProductsRoutes    *Router
	HealthcheckRoutes *Router
}

func NewRoutes(controllers *controller.Controllers) *Routes {
	return &Routes{
		ClientsRoutes:     Clients(controllers.ClientsController),
		OrdersRoutes:      Orders(controllers.OrdersController),
		ProductsRoutes:    Products(controllers.ProductsController),
		HealthcheckRoutes: HealthCheck(),
	}
}
