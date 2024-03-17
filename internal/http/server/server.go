package server

import (
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/raphael-foliveira/chi-gorm/internal/config"
	"github.com/raphael-foliveira/chi-gorm/internal/database"
	"github.com/raphael-foliveira/chi-gorm/internal/http/controller"
	"github.com/raphael-foliveira/chi-gorm/internal/http/routes"
	"github.com/raphael-foliveira/chi-gorm/internal/repository"
	"github.com/raphael-foliveira/chi-gorm/internal/service"
)

type app struct {
	Db           *database.DB
	Dependencies *ServerDependencies
}

func NewApp(db *database.DB) *app {
	return &app{Db: db, Dependencies: injectDependencies(db)}
}

func (a *app) Start() error {
	app := a.CreateMainRouter()
	slog.Info("listening on port 3000")
	return http.ListenAndServe(":3000", app)
}

func (a *app) CreateMainRouter() *chi.Mux {
	mainRouter := chi.NewRouter()
	a.attachMiddleware(mainRouter)
	a.mountRoutes(mainRouter)
	return mainRouter
}

func (a *app) mountRoutes(r *chi.Mux) {
	r.Mount("/clients", a.Dependencies.Routes.ClientsRoutes)
	r.Mount("/products", a.Dependencies.Routes.ProductsRoutes)
	r.Mount("/orders", a.Dependencies.Routes.OrdersRoutes)
	r.Mount("/health-check", a.Dependencies.Routes.HealthcheckRoutes)
}

func (a *app) attachMiddleware(r *chi.Mux) {
	r.Use(middleware.Logger)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"*"},
	}))
}

type ServerDependencies struct {
	Repositories *repository.Repositories
	Services     *service.Services
	Controllers  *controller.Controllers
	Routes       *routes.Routes
}

func injectDependencies(db *database.DB) *ServerDependencies {
	repositories := repository.NewRepositories(db)
	services := service.NewServices(repositories, config.Config())
	controllers := controller.NewControllers(services)
	routes := routes.NewRoutes(controllers)
	return &ServerDependencies{
		Repositories: repositories,
		Services:     services,
		Controllers:  controllers,
		Routes:       routes,
	}
}
