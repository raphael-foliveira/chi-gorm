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
	Db           *gorm.DB
	Repositories *repository.Repositories
	Services     *service.Services
	Controllers  *controller.Controllers
	Routers      *routes.Routers
}

func NewApp(db *gorm.DB) *App {
	return &App{Db: db}
}

func (a *App) Start() error {
	app := a.CreateMainRouter()
	slog.Info("listening on port 3000")
	return http.ListenAndServe(":3000", app)
}

func (a *App) CreateMainRouter() *chi.Mux {
	mainRouter := chi.NewRouter()
	a.attachMiddleware(mainRouter)
	a.injectDependencies()
	a.mountRoutes(mainRouter, a.Routers)
	return mainRouter
}

func (a *App) injectDependencies() {
	a.Repositories = repository.NewRepositories(a.Db)
	a.Services = service.NewServices(a.Repositories)
	a.Controllers = controller.NewControllers(a.Services)
	a.Routers = routes.NewRouters(a.Controllers)
}

func (a *App) mountRoutes(r *chi.Mux, rts *routes.Routers) {
	r.Mount("/clients", rts.Clients)
	r.Mount("/products", rts.Products)
	r.Mount("/orders", rts.Orders)
	r.Get("/health-check", rts.HealthCheck)
}

func (a *App) attachMiddleware(r *chi.Mux) {
	r.Use(middleware.Logger)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"*"},
	}))
}
