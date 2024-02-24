package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/raphael-foliveira/chi-gorm/internal/http/controller"
)

func Clients() *chi.Mux {
	c := controller.Clients()
	router := chi.NewRouter()
	router.Get("/", useHandler(c.List))
	router.Get("/{id}", useHandler(c.Get))
	router.Get("/{id}/products", useHandler(c.GetProducts))
	router.Post("/", useHandler(c.Create))
	router.Delete("/{id}", useHandler(c.Delete))
	router.Put("/{id}", useHandler(c.Update))
	return router
}
