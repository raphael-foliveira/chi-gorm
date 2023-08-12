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

func mountRouters(r *chi.Mux, db *db.DB) {
	clientsRouter, err := client.NewRouter(db)
	if err != nil {
		panic(err)
	}
	productsRouter, err := product.NewRouter(db)
	if err != nil {
		panic(err)
	}
	ordersRouter, err := order.NewRouter(db)
	if err != nil {
		panic(err)
	}
	r.Mount("/clients", clientsRouter)
	r.Mount("/products", productsRouter)
	r.Mount("/orders", ordersRouter)
}

func Start(db *db.DB) error {
	mainRouter := chi.NewRouter()
	attachMiddleware(mainRouter)
	mainRouter.Get("/", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]string{"message": "Hello World!"})
	})
	mountRouters(mainRouter, db)
	fmt.Println("listening on port 3000")
	return http.ListenAndServe(":3000", mainRouter)
}