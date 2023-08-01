package client

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/raphael-foliveira/chi-gorm/internal/db"
)

func AttachRouter(r *chi.Mux, db *db.DB) {
	err := db.AutoMigrate(&Client{})
	if err != nil {
		panic(err)
	}
	repository := NewRepository(db)
	controller := NewController(repository)

	clientRouter := chi.NewRouter()
	clientRouter.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			next.ServeHTTP(w, r)
		})
	})

	clientRouter.Get("/", controller.List)
	clientRouter.Post("/", controller.Create)
	clientRouter.Get("/{id}", controller.Get)
	clientRouter.Delete("/{id}", controller.Delete)
	clientRouter.Put("/{id}", controller.Update)

	r.Mount("/clients", clientRouter)
}
