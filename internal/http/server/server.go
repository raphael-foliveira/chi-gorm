package server

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/raphael-foliveira/chi-gorm/internal/database"
	"github.com/raphael-foliveira/chi-gorm/internal/http/controllers"
	"github.com/raphael-foliveira/chi-gorm/internal/http/routes"
	"github.com/raphael-foliveira/chi-gorm/internal/repository"
	"github.com/raphael-foliveira/chi-gorm/internal/services"
)

func Start() error {
	app := CreateApp()
	fmt.Println("listening on port 3000")
	return http.ListenAndServe(":3000", app)
}

func CreateApp() *chi.Mux {
	mainRouter := chi.NewRouter()
	attachMiddleware(mainRouter)
	injectDependencies(mainRouter)
	return mainRouter
}

func injectDependencies(r *chi.Mux) {
	db := database.GetDb()
	clientsRepo := repository.NewClients(db)
	productsRepo := repository.NewProducts(db)
	ordersRepo := repository.NewOrders(db)

	clientsService := services.NewClients(clientsRepo)
	productsService := services.NewProducts(productsRepo)
	ordersService := services.NewOrders(ordersRepo)

	clientsController := controllers.NewClients(clientsService)
	productsController := controllers.NewProducts(productsService)
	ordersController := controllers.NewOrders(ordersService)

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
