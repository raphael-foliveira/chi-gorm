package controller

import (
	"github.com/go-chi/chi/v5"
)

func clientsRoutes() *chi.Mux {
	router := chi.NewRouter()
	router.Get("/", useHandler(Clients.List))
	router.Get("/{id}", useHandler(Clients.Get))
	router.Get("/{id}/products", useHandler(Clients.GetProducts))
	router.Post("/", useHandler(Clients.Create))
	router.Delete("/{id}", useHandler(Clients.Delete))
	router.Put("/{id}", useHandler(Clients.Update))

	return router
}

func productsRoutes() *chi.Mux {
	router := chi.NewRouter()
	router.Get("/", useHandler(Products.List))
	router.Post("/", useHandler(Products.Create))
	router.Get("/{id}", useHandler(Products.Get))
	router.Delete("/{id}", useHandler(Products.Delete))
	router.Put("/{id}", useHandler(Products.Update))

	return router
}

func ordersRoutes() *chi.Mux {
	router := chi.NewRouter()
	router.Get("/", useHandler(Orders.List))
	router.Post("/", useHandler(Orders.Create))
	router.Get("/{id}", useHandler(Orders.Get))
	router.Delete("/{id}", useHandler(Orders.Delete))
	router.Put("/{id}", useHandler(Orders.Update))

	return router
}

func healthCheckRoutes() *chi.Mux {
	router := chi.NewRouter()
	router.Get("/", useHandler(HealthCheck.healthCheck))

	return router
}

func Mount(mux *chi.Mux) {
	mux.Mount("/clients", clientsRoutes())
	mux.Mount("/products", productsRoutes())
	mux.Mount("/orders", ordersRoutes())
	mux.Mount("/health-check", healthCheckRoutes())
}
