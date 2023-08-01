package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/raphael-foliveira/chi-gorm/internal/db"
	"github.com/raphael-foliveira/chi-gorm/internal/modules/client"
	"github.com/raphael-foliveira/chi-gorm/internal/modules/order"
	"github.com/raphael-foliveira/chi-gorm/internal/modules/product"
)

func attachMiddleware(r *chi.Mux) {
	r.Use(middleware.Logger)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"*"},
	}))
}

func attachRoutes(r *chi.Mux, db *db.DB) {
	client.AttachRouter(r, db)
	product.AttachRouter(r, db)
	order.AttachRouter(r, db)
}

func Start(db *db.DB) error {
	mainRouter := chi.NewRouter()
	attachMiddleware(mainRouter)
	mainRouter.Get("/", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]string{"message": "Hello World!"})
	})
	attachRoutes(mainRouter, db)
	fmt.Println("listening on port 3000")
	return http.ListenAndServe(":3000", mainRouter)
}
