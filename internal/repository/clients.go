package repository

import (
	"github.com/raphael-foliveira/chi-gorm/internal/database"
	"github.com/raphael-foliveira/chi-gorm/internal/entities"
)

type ClientsRepository interface {
	Repository[entities.Client]
}

type Clients struct {
	*repository[entities.Client]
}

func NewClients(db *database.DB) *Clients {
	return &Clients{newRepository[entities.Client](db)}
}
