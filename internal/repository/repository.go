package repository

import (
	"github.com/raphael-foliveira/chi-gorm/internal/database"
	"github.com/raphael-foliveira/chi-gorm/internal/ports"
	"gorm.io/gorm"
)

var _ ports.Repository[any] = &Repository[any]{}

type Repository[T any] struct {
	db *gorm.DB
}

func NewRepository[T any](db *gorm.DB) *Repository[T] {
	return &Repository[T]{
		db: db,
	}
}

func (r *Repository[T]) List(conds ...any) ([]T, error) {
	entities := []T{}
	return entities, database.DB.Find(&entities, conds...).Error
}

func (r *Repository[T]) Get(id uint) (*T, error) {
	entity := new(T)
	return entity, database.DB.Model(new(T)).First(entity, id).Error
}

func (r *Repository[T]) Create(entity *T) error {
	return database.DB.Create(entity).Error
}

func (r *Repository[T]) Update(entity *T) error {
	return database.DB.Save(entity).Error
}

func (r *Repository[T]) Delete(id uint) error {
	entity := new(T)
	return database.DB.Delete(entity, id).Error
}
