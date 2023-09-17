package server

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/raphael-foliveira/chi-gorm/pkg/http/controllers"
	"github.com/raphael-foliveira/chi-gorm/pkg/http/routes"
	"github.com/raphael-foliveira/chi-gorm/pkg/persistence/db"
	"github.com/raphael-foliveira/chi-gorm/pkg/persistence/store"
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
	clientsStore := store.NewClients(db.Db)
	productsStore := store.NewProducts(db.Db)
	ordersStore := store.NewOrders(db.Db)

	clientsController := controllers.NewClients(clientsStore)
	productsController := controllers.NewProducts(productsStore)
	ordersController := controllers.NewOrders(ordersStore)

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
