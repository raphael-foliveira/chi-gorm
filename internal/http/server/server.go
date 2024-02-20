package server

import (
	"log/slog"
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

type App struct {
	db *gorm.DB
}

func NewApp(db *gorm.DB) *App {
	return &App{db}
}

func (s *App) Start() error {
	app := s.CreateRouter()
	slog.Info("listening on port 3000")
	return http.ListenAndServe(":3000", app)
}

func (s *App) CreateRouter() *chi.Mux {
	mainRouter := chi.NewRouter()
	s.attachMiddleware(mainRouter)
	controllers := s.injectDependencies()
	s.mountRoutes(mainRouter, controllers)
	return mainRouter
}

func (s *App) injectDependencies() *routes.Routers {
	repositories := repository.NewRepositories(s.db)
	services := service.NewServices(repositories)
	controllers := controller.NewControllers(services)
	routes := routes.NewRouters(controllers)
	return routes
}

func (s *App) mountRoutes(r *chi.Mux, rts *routes.Routers) {
	r.Mount("/clients", rts.Clients)
	r.Mount("/products", rts.Products)
	r.Mount("/orders", rts.Orders)
	r.Get("/health-check", rts.HealthCheck)
}

func (s *App) attachMiddleware(r *chi.Mux) {
	r.Use(middleware.Logger)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"*"},
	}))
}
