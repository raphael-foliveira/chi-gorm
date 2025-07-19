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

	// attach middleware
	r.Use(middleware.Logger)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"*"},
	}))

	return &Server{
		Mux: r,
	}
}

func (s *Server) Mount() {
	clientsController := NewClientsController(s.ClientsRepository, s.OrdersRepository)
	healthCheckController := NewHealthCheckController()
	ordersController := NewOrdersController(s.OrdersRepository)
	productsController := NewProductsController(s.ProductsRepository)

	clientsRouter := chi.NewRouter()
	clientsRouter.Get("/", useHandler(clientsController.List))
	clientsRouter.Get("/{id}", useHandler(clientsController.Get))
	clientsRouter.Get("/{id}/products", useHandler(clientsController.GetProducts))
	clientsRouter.Post("/", useHandler(clientsController.Create))
	clientsRouter.Delete("/{id}", useHandler(clientsController.Delete))
	clientsRouter.Put("/{id}", useHandler(clientsController.Update))
	s.Mux.Mount("/clients", clientsRouter)

	healthCheckRouter := chi.NewRouter()
	healthCheckRouter.Get("/", useHandler(healthCheckController.healthCheck))
	s.Mux.Mount("/health-check", healthCheckRouter)

	ordersRouter := chi.NewRouter()
	ordersRouter.Get("/", useHandler(ordersController.List))
	ordersRouter.Get("/{id}", useHandler(ordersController.Get))
	ordersRouter.Post("/", useHandler(ordersController.Create))
	ordersRouter.Delete("/{id}", useHandler(ordersController.Delete))
	ordersRouter.Put("/{id}", useHandler(ordersController.Update))
	s.Mux.Mount("/orders", ordersRouter)

	productsRouter := chi.NewRouter()
	productsRouter.Get("/", useHandler(productsController.List))
	productsRouter.Get("/{id}", useHandler(productsController.Get))
	productsRouter.Post("/", useHandler(productsController.Create))
	productsRouter.Delete("/{id}", useHandler(productsController.Delete))
	productsRouter.Put("/{id}", useHandler(productsController.Update))
	s.Mux.Mount("/products", productsRouter)
}
