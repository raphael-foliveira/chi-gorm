package repository

import (
	"github.com/raphael-foliveira/chi-gorm/internal/entities"
	"gorm.io/gorm"
)

type Clients Repository[entities.Client]

type clients struct {
	*repository[entities.Client]
}

func NewClients(db *gorm.DB) Clients {
	db.AutoMigrate(&entities.Client{})
	return &clients{newRepository[entities.Client](db)}
}
