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
	mw "github.com/raphael-foliveira/chi-gorm/pkg/middleware"
)

func attachMiddleware(r *chi.Mux) {
	r.Use(middleware.Logger)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"*"},
	}))
	r.Use(mw.Json)
}

func mountRouters(r *chi.Mux, db *db.DB) {
	clientsRouter := client.Init(db)
	productsRouter := product.Init(db)
	ordersRouter := order.Init(db)
	r.Mount("/clients", clientsRouter)
	r.Mount("/products", productsRouter)
	r.Mount("/orders", ordersRouter)
}

func Start(db *db.DB) error {
	db.AutoMigrate(&client.Client{}, &product.Product{}, &order.Order{})
	mainRouter := chi.NewRouter()
	attachMiddleware(mainRouter)
	mainRouter.Get("/", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]string{"message": "Hello World!"})
	})
	mountRouters(mainRouter, db)
	fmt.Println("listening on port 3000")
	return http.ListenAndServe(":3000", mainRouter)
}
