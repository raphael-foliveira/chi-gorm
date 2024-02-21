package repository

import (
	"github.com/raphael-foliveira/chi-gorm/internal/entities"
	"gorm.io/gorm"
)

type ClientsRepository interface {
	Repository[entities.Client]
}

type clients struct {
	*repository[entities.Client]
}

func NewClients(db *gorm.DB) *clients {
	return &clients{&repository[entities.Client]{db}}
}
