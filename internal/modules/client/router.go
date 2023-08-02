package client

import (
	"github.com/go-chi/chi/v5"
	"github.com/raphael-foliveira/chi-gorm/internal/db"
	"github.com/raphael-foliveira/chi-gorm/pkg/middleware"
)

func MountRouter(r *chi.Mux, db *db.DB) {
	err := db.AutoMigrate(&Client{})
	if err != nil {
		panic(err)
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

	r.Mount("/clients", clientRouter)
}
