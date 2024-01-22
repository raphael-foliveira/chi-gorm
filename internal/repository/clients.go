package repository

import (
	"github.com/raphael-foliveira/chi-gorm/internal/entities"
	"gorm.io/gorm"
)

type clients struct {
	*Repository[entities.Client]
}

func NewClients(db *gorm.DB) *clients {
	return &clients{&Repository[entities.Client]{db}}
}

func (c *clients) Delete(entity *entities.Client) error {
	return c.db.Delete(entity).Error
}
