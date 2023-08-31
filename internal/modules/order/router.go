package order

import (
	"github.com/go-chi/chi/v5"
	"github.com/raphael-foliveira/chi-gorm/pkg/middleware"
)

func NewRouter(controller *Controller) *chi.Mux {

	ordersRouter := chi.NewRouter()
	ordersRouter.Use(middleware.Json)

	ordersRouter.Get("/", controller.List)
	ordersRouter.Post("/", controller.Create)
	ordersRouter.Get("/{id}", controller.Get)
	ordersRouter.Delete("/{id}", controller.Delete)
	ordersRouter.Put("/{id}", controller.Update)
	return ordersRouter
}
