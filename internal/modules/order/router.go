package order

import (
	"github.com/go-chi/chi/v5"
	"github.com/raphael-foliveira/chi-gorm/pkg/handler"
)

func NewRouter(controller *Controller) *chi.Mux {
	router := chi.NewRouter()
	router.Get("/", handler.Wrap(controller.List))
	router.Post("/", handler.Wrap(controller.Create))
	router.Get("/{id}", handler.Wrap(controller.Get))
	router.Delete("/{id}", handler.Wrap(controller.Delete))
	router.Put("/{id}", handler.Wrap(controller.Update))
	return router
}
