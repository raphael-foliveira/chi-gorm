package repository

import (
	"github.com/raphael-foliveira/chi-gorm/internal/entities"
)

type clients struct {
	*repository[entities.Client]
}

func NewClients() *clients {
	return &clients{newRepository[entities.Client]()}
}
