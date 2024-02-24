package repository

import (
	"github.com/raphael-foliveira/chi-gorm/internal/database"
	"github.com/raphael-foliveira/chi-gorm/internal/entities"
	"gorm.io/gorm"
)

type ClientsRepository interface {
	Repository[entities.Client]
}

func Clients() *clients {
	return NewClients(database.Db())
}

type clients struct {
	*repository[entities.Client]
}

func NewClients(db *gorm.DB) *clients {
	return &clients{newRepository[entities.Client](db)}
}
