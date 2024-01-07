package server

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/raphael-foliveira/chi-gorm/internal/http/controller"
	"github.com/raphael-foliveira/chi-gorm/internal/http/routes"
	"github.com/raphael-foliveira/chi-gorm/internal/repository"
	"github.com/raphael-foliveira/chi-gorm/internal/service"
	"gorm.io/gorm"
)

type Server struct {
	db *gorm.DB
}

func NewServer(db *gorm.DB) *Server {
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
	controllers := s.injectDependencies()
	s.mountRoutes(mainRouter, controllers)
	return mainRouter
}

func (s *Server) injectDependencies() *controller.Controllers {
	repositories := repository.NewRepositories(s.db)
	services := service.NewServices(repositories)
	controllers := controller.NewControllers(services)
	return controllers
}

func (s *Server) mountRoutes(r *chi.Mux, c *controller.Controllers) {
	clientsRoutes := routes.Clients(c.Clients)
	productsRoutes := routes.Products(c.Products)
	ordersRoutes := routes.Orders(c.Orders)

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
