package repository

import (
	"github.com/raphael-foliveira/chi-gorm/internal/entities"
	"gorm.io/gorm"
)

type Clients interface {
	Repository[entities.Client]
}

type clients struct {
	*repository[entities.Client]
}

func NewClients(db *gorm.DB) Clients {
	return &clients{&repository[entities.Client]{db}}
}

func (c *clients) Delete(entity *entities.Client) error {
	return c.db.Delete(entity).Error
}
