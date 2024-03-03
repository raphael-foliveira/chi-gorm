package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/raphael-foliveira/chi-gorm/internal/http/controller"
	"github.com/raphael-foliveira/chi-gorm/internal/http/middleware"
)

func Orders() *chi.Mux {
	controller := controller.Orders()
	authMw := middleware.Auth()
	router := chi.NewRouter()
	router.Get("/", useHandler(controller.List))
	router.Post("/", useHandler(authMw.CheckToken(controller.Create)))
	router.Get("/{id}", useHandler(controller.Get))
	router.Delete("/{id}", useHandler(controller.Delete))
	router.Put("/{id}", useHandler(controller.Update))
	return router
}
