package order

import (
	"github.com/go-chi/chi/v5"
	"github.com/raphael-foliveira/chi-gorm/internal/db"
	"github.com/raphael-foliveira/chi-gorm/pkg/middleware"
)



func AttachRouter(r *chi.Mux, db *db.DB) {
	err := db.AutoMigrate(&Order{})
	if err != nil {
		panic(err)
	}
	repository := NewRepository(db)
	controller := NewController(repository)

	orderRouter := chi.NewRouter()
	orderRouter.Use(middleware.Json)

	orderRouter.Get("/", controller.List)
	orderRouter.Post("/", controller.Create)
	orderRouter.Get("/{id}", controller.Get)
	orderRouter.Delete("/{id}", controller.Delete)
	orderRouter.Put("/{id}", controller.Update)

	r.Mount("/orders", orderRouter)
}
