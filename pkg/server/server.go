package server

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/raphael-foliveira/chi-gorm/pkg/controllers"
	"github.com/raphael-foliveira/chi-gorm/pkg/db"
	"github.com/raphael-foliveira/chi-gorm/pkg/models"
	"github.com/raphael-foliveira/chi-gorm/pkg/repositories"
	"github.com/raphael-foliveira/chi-gorm/pkg/routes"
)

func Start(db *db.DB) error {
	db.AutoMigrate(&models.Client{}, &models.Product{}, &models.Order{})
	mainRouter := chi.NewRouter()
	attachMiddleware(mainRouter)
	injectDependencies(mainRouter, db)
	fmt.Println("listening on port 3000")
	return http.ListenAndServe(":3000", mainRouter)
}

func injectDependencies(r *chi.Mux, db *db.DB) {
	clientsRepository := repositories.NewClient(db)
	productsRepository := repositories.NewProducts(db)
	ordersRepository := repositories.NewOrders(db)

	clientsController := controllers.NewClients(clientsRepository)
	productsController := controllers.NewProducts(productsRepository)
	ordersController := controllers.NewOrders(ordersRepository)

	clientsRoutes := routes.Clients(clientsController)
	productsRoutes := routes.Products(productsController)
	ordersRoutes := routes.Orders(ordersController)

	r.Mount("/clients", clientsRoutes)
	r.Mount("/products", productsRoutes)
	r.Mount("/orders", ordersRoutes)
}

func attachMiddleware(r *chi.Mux) {
	r.Use(middleware.Logger)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"*"},
	}))
}
