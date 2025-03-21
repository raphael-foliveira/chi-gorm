package repository

import (
	"github.com/raphael-foliveira/chi-gorm/internal/entities"
	"github.com/raphael-foliveira/chi-gorm/internal/ports"
	"gorm.io/gorm"
)

var _ ports.ClientsRepository = &Clients{}

type Clients struct {
	*Repository[entities.Client]
}

func NewClients(db *gorm.DB) *Clients {
	return &Clients{NewRepository[entities.Client](db)}
}
