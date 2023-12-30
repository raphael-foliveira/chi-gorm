package repository

import (
	"github.com/raphael-foliveira/chi-gorm/internal/database"
	"github.com/raphael-foliveira/chi-gorm/internal/entities"
)

var Clients ClientsRepository = &clients{&repository[entities.Client]{}}

type ClientsRepository interface {
	Repository[entities.Client]
}

type clients struct {
	*repository[entities.Client]
}

func (c *clients) Delete(entity *entities.Client) error {
	return database.Db.Delete(entity).Error
}
