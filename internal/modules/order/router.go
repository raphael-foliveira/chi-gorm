package order

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/raphael-foliveira/chi-gorm/internal/db"
)

func AttachRouter(r *chi.Mux, db *db.DB) {
	err := db.AutoMigrate(&Order{})
	if err != nil {
		panic(err)
	}
	repository := NewRepository(db)
	controller := NewController(repository)

	orderRouter := chi.NewRouter()
	orderRouter.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			next.ServeHTTP(w, r)
		})
	})

	orderRouter.Get("/", controller.List)
	orderRouter.Post("/", controller.Create)
	orderRouter.Get("/{id}", controller.Get)
	orderRouter.Delete("/{id}", controller.Delete)
	orderRouter.Put("/{id}", controller.Update)

	r.Mount("/orders", orderRouter)
}
