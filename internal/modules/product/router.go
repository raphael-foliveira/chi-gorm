package product

import (
	"github.com/go-chi/chi/v5"
	"github.com/raphael-foliveira/chi-gorm/pkg/middleware"
)

func NewRouter(controller *Controller) *chi.Mux {
	productRouter := chi.NewRouter()
	productRouter.Use(middleware.Json)
	productRouter.Get("/", controller.List)
	productRouter.Post("/", controller.Create)
	productRouter.Get("/{id}", controller.Get)
	productRouter.Delete("/{id}", controller.Delete)
	productRouter.Put("/{id}", controller.Update)
	return productRouter
}
