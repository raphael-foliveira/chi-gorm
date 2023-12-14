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
	"gorm.io/gorm"
)

type Server struct {
	Db *gorm.DB
}

func NewServer(dialector gorm.Dialector) *Server {
	db := database.InitDb(dialector)
	return &Server{db}
}

func (s *Server) Start() error {
	app := s.CreateApp()
	fmt.Println("listening on port 3000")
	return http.ListenAndServe(":3000", app)
}

func (s *Server) CreateApp() *chi.Mux {
	mainRouter := chi.NewRouter()
	s.attachMiddleware(mainRouter)
	s.injectDependencies(mainRouter)
	return mainRouter
}

func (s *Server) injectDependencies(r *chi.Mux) {
	clientsRepo := repository.NewClients(s.Db)
	productsRepo := repository.NewProducts(s.Db)
	ordersRepo := repository.NewOrders(s.Db)

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

func (s *Server) attachMiddleware(r *chi.Mux) {
	r.Use(middleware.Logger)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"*"},
	}))
}
