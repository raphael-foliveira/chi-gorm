package product

import (
	"github.com/go-chi/chi/v5"
	"github.com/raphael-foliveira/chi-gorm/internal/db"
	"github.com/raphael-foliveira/chi-gorm/pkg/middleware"
)

func MountRouter(r *chi.Mux, db *db.DB) {
	err := db.AutoMigrate(&Product{})
	if err != nil {
		panic(err)
	}
	repository := NewRepository(db)
	controller := NewController(repository)

	productRouter := chi.NewRouter()
	productRouter.Use(middleware.Json)

	productRouter.Get("/", controller.List)
	productRouter.Post("/", controller.Create)
	productRouter.Get("/{id}", controller.Get)
	productRouter.Delete("/{id}", controller.Delete)
	productRouter.Put("/{id}", controller.Update)

	r.Mount("/products", productRouter)
}
