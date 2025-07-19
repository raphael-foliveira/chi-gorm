package api

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/raphael-foliveira/chi-gorm/internal/ports"
)

type Server struct {
	*chi.Mux
	ClientsRepository  ports.ClientsRepository
	JwtService         ports.JwtService
	OrdersRepository   ports.OrdersRepository
	ProductsRepository ports.ProductsRepository
}

func NewServer() *Server {
	r := chi.NewRouter()
	attachMiddleware(r)
	return &Server{
		Mux: r,
	}
}

func (s *Server) Mount() {
	clientsController := NewClientsController(s.ClientsRepository, s.OrdersRepository)
	healthCheckController := NewHealthCheckController()
	ordersController := NewOrdersController(s.OrdersRepository)
	productsController := NewProductsController(s.ProductsRepository)

	clientsRouter := NewRouter()
	clientsRouter.Get("/", clientsController.List)
	clientsRouter.Get("/{id}", clientsController.Get)
	clientsRouter.Get("/{id}/products", clientsController.GetProducts)
	clientsRouter.Post("/", clientsController.Create)
	clientsRouter.Delete("/{id}", clientsController.Delete)
	clientsRouter.Put("/{id}", clientsController.Update)

	healthCheckRouter := NewRouter()
	healthCheckRouter.Get("/", healthCheckController.healthCheck)

	ordersRouter := NewRouter()
	ordersRouter.Get("/", ordersController.List)
	ordersRouter.Get("/{id}", ordersController.Get)
	ordersRouter.Post("/", ordersController.Create)
	ordersRouter.Delete("/{id}", ordersController.Delete)
	ordersRouter.Put("/{id}", ordersController.Update)

	productsRouter := NewRouter()
	productsRouter.Get("/", productsController.List)
	productsRouter.Get("/{id}", productsController.Get)
	productsRouter.Post("/", productsController.Create)
	productsRouter.Delete("/{id}", productsController.Delete)
	productsRouter.Put("/{id}", productsController.Update)

	s.Mux.Mount("/clients", clientsRouter)
	s.Mux.Mount("/health-check", healthCheckRouter)
	s.Mux.Mount("/orders", ordersRouter)
	s.Mux.Mount("/products", productsRouter)
}

func attachMiddleware(r *chi.Mux) {
	r.Use(middleware.Logger)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"*"},
	}))
}
