package product

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/raphael-foliveira/chi-gorm/internal/db"
)

func AttachRouter(r *chi.Mux, db *db.DB) {
	err := db.AutoMigrate(&Product{})
	if err != nil {
		panic(err)
	}
	repository := NewRepository(db)
	controller := NewController(repository)

	productRouter := chi.NewRouter()
	productRouter.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			next.ServeHTTP(w, r)
		})
	})
	productRouter.Get("/", controller.List)
	productRouter.Post("/", controller.Create)
	productRouter.Get("/{id}", controller.Get)
	productRouter.Delete("/{id}", controller.Delete)
	productRouter.Put("/{id}", controller.Update)

	r.Mount("/products", productRouter)
}
