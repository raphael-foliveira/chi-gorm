package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/raphael-foliveira/chi-gorm/internal/http/controller"
)

func Clients(c *controller.Clients) *Router {
	router := &Router{chi.NewRouter()}
	router.Get("/", c.List)
	router.Get("/{id}", c.Get)
	router.Get("/{id}/products", c.GetProducts)
	router.Post("/", c.Create)
	router.Delete("/{id}", c.Delete)
	router.Put("/{id}", c.Update)
	return router
}
