package server

import (
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/raphael-foliveira/chi-gorm/internal/database"
	"github.com/raphael-foliveira/chi-gorm/internal/http/routes"
	"gorm.io/gorm"
)

type app struct {
	Db *gorm.DB
}

func NewApp(db *gorm.DB) *app {
	return &app{Db: db}
}

func App() *app {
	return NewApp(database.Db())
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
	r.Mount("/clients", routes.Clients())
	r.Mount("/products", routes.Products())
	r.Mount("/orders", routes.Orders())
	r.Get("/health-check", routes.HealthCheck())
}

func (a *app) attachMiddleware(r *chi.Mux) {
	r.Use(middleware.Logger)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"*"},
	}))
}
