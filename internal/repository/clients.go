package repository

import (
	"github.com/raphael-foliveira/chi-gorm/internal/entities"
)

type Clients struct {
	*repository[entities.Client]
}

func NewClients() *Clients {
	return &Clients{newRepository[entities.Client]()}
}
