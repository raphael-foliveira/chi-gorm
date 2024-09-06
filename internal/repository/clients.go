package repository

import (
	"github.com/raphael-foliveira/chi-gorm/internal/entities"
)

type Clients struct {
	*Repository[entities.Client]
}

func NewClients() *Clients {
	return &Clients{New[entities.Client]()}
}
