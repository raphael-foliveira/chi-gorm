package server

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/raphael-foliveira/chi-gorm/internal/ports"
)

type Server struct {
	*chi.Mux
}

func New() *Server {
	r := chi.NewRouter()
	attachMiddleware(r)
	return &Server{
		Mux: r,
	}
}

func (s *Server) Mount(c ...ports.Controller) {
	for _, cc := range c {
		cc.Mount(s.Mux)
	}
}

func attachMiddleware(r *chi.Mux) {
	r.Use(middleware.Logger)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"*"},
	}))
}
