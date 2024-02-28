package repository

import (
	"github.com/raphael-foliveira/chi-gorm/internal/entities"
	"gorm.io/gorm"
)

type clientsRepository struct {
	*repository[entities.Client]
}

func NewClientsRepository(db *gorm.DB) *clientsRepository {
	return &clientsRepository{newRepository[entities.Client](db)}
}
