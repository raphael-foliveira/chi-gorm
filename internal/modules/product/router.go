package product

import (
	"github.com/go-chi/chi/v5"
)

func NewRouter(controller *Controller) *chi.Mux {
	router := chi.NewRouter()
	router.Get("/", controller.List)
	router.Post("/", controller.Create)
	router.Get("/{id}", controller.Get)
	router.Delete("/{id}", controller.Delete)
	router.Put("/{id}", controller.Update)
	return router
}
