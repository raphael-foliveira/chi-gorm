package client

import (
	"github.com/go-chi/chi/v5"
	"github.com/raphael-foliveira/chi-gorm/pkg/middleware"
)

func NewRouter(controller *Controller) *chi.Mux {

	clientRouter := chi.NewRouter()
	clientRouter.Use(middleware.Json)

	clientRouter.Get("/", controller.List)
	clientRouter.Post("/", controller.Create)
	clientRouter.Get("/{id}", controller.Get)
	clientRouter.Delete("/{id}", controller.Delete)
	clientRouter.Put("/{id}", controller.Update)

	return clientRouter
}
