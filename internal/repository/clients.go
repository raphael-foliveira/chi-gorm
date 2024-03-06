package repository

import (
	"github.com/raphael-foliveira/chi-gorm/internal/entities"
	"gorm.io/gorm"
)

type ClientsRepository interface {
	Repository[entities.Client]
}

type Clients struct {
	*repository[entities.Client]
}

func NewClients(db *gorm.DB) *Clients {
	return &Clients{newRepository[entities.Client](db)}
}
