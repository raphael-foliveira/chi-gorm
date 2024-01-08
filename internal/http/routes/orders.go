package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/raphael-foliveira/chi-gorm/internal/http/controller"
)

func Orders(c *controller.Orders) *chi.Mux {
	router := chi.NewRouter()
	router.Get("/", wrap(c.List))
	router.Post("/", wrap(c.Create))
	router.Get("/{id}", wrap(c.Get))
	router.Delete("/{id}", wrap(c.Delete))
	router.Put("/{id}", wrap(c.Update))
	return router
}
