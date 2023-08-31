package client

import (
	"github.com/go-chi/chi/v5"
	"github.com/raphael-foliveira/chi-gorm/internal/db"
)

func Init(db *db.DB) *chi.Mux {
	repository := NewRepository(db)
	controller := NewController(repository)
	return NewRouter(controller)
}
