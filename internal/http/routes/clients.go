package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/raphael-foliveira/chi-gorm/internal/http/controller"
)

func Clients() *chi.Mux {
	router := chi.NewRouter()
	router.Get("/", wrap(controller.Clients.List))
	router.Post("/", wrap(controller.Clients.Create))
	router.Get("/{id}", wrap(controller.Clients.Get))
	router.Delete("/{id}", wrap(controller.Clients.Delete))
	router.Put("/{id}", wrap(controller.Clients.Update))
	return router
}
