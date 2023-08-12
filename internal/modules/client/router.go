package client

import (
	"github.com/go-chi/chi/v5"
	"github.com/raphael-foliveira/chi-gorm/internal/db"
	"github.com/raphael-foliveira/chi-gorm/pkg/middleware"
)

func NewRouter(db *db.DB) (*chi.Mux, error) {
	err := db.AutoMigrate(&Client{})
	if err != nil {
		return nil, err
	}
	repository := NewRepository(db)
	controller := NewController(repository)

	clientRouter := chi.NewRouter()
	clientRouter.Use(middleware.Json)

	clientRouter.Get("/", controller.List)
	clientRouter.Post("/", controller.Create)
	clientRouter.Get("/{id}", controller.Get)
	clientRouter.Delete("/{id}", controller.Delete)
	clientRouter.Put("/{id}", controller.Update)

	return clientRouter, nil
}