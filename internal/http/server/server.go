package server

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/raphael-foliveira/chi-gorm/internal/http/routes"
)

type Server struct{}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Start() error {
	app := s.CreateApp()
	fmt.Println("listening on port 3000")
	return http.ListenAndServe(":3000", app)
}

func (s *Server) CreateApp() *chi.Mux {
	mainRouter := chi.NewRouter()
	s.attachMiddleware(mainRouter)
	s.mountRoutes(mainRouter)
	return mainRouter
}

func (s *Server) mountRoutes(r *chi.Mux) {
	clientsRoutes := routes.Clients()
	productsRoutes := routes.Products()
	ordersRoutes := routes.Orders()

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
