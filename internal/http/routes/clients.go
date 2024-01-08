package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/raphael-foliveira/chi-gorm/internal/http/controller"
)

func Clients(c *controller.Clients) *chi.Mux {
	router := chi.NewRouter()
	router.Get("/", wrap(c.List))
	router.Get("/{id}", wrap(c.Get))
	router.Get("/{id}/products", wrap(c.GetProducts))
	router.Post("/", wrap(c.Create))
	router.Delete("/{id}", wrap(c.Delete))
	router.Put("/{id}", wrap(c.Update))
	return router
}
