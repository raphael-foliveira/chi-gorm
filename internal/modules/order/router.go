package order

import (
	"github.com/go-chi/chi/v5"
	"github.com/raphael-foliveira/chi-gorm/internal/db"
	"github.com/raphael-foliveira/chi-gorm/pkg/middleware"
)

func NewRouter(db *db.DB) (*chi.Mux, error) {
	err := db.AutoMigrate(&Order{})
	if err != nil {
		return nil, err
	}
	repository := NewRepository(db)
	controller := NewController(repository)

	ordersRouter := chi.NewRouter()
	ordersRouter.Use(middleware.Json)

	ordersRouter.Get("/", controller.List)
	ordersRouter.Post("/", controller.Create)
	ordersRouter.Get("/{id}", controller.Get)
	ordersRouter.Delete("/{id}", controller.Delete)
	ordersRouter.Put("/{id}", controller.Update)
	return ordersRouter, nil
}
