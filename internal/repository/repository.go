package repository

import (
	"github.com/raphael-foliveira/chi-gorm/internal/database"
)

type Repository[T any] interface {
	List(...any) ([]T, error)
	Get(uint) (*T, error)
	Create(*T) error
	Update(*T) error
	Delete(*T) error
}

type repository[T any] struct {
	db *database.DB
}

func newRepository[T any](db *database.DB) *repository[T] {
	return &repository[T]{db}
}

func (r *repository[T]) List(conds ...any) ([]T, error) {
	entities := []T{}
	return entities, r.db.Find(&entities, conds...).Error
}

func (r *repository[T]) Get(id uint) (*T, error) {
	entity := new(T)
	return entity, r.db.Model(new(T)).First(entity, id).Error
}

func (r *repository[T]) Create(entity *T) error {
	return r.db.Create(entity).Error
}

func (r *repository[T]) Update(entity *T) error {
	return r.db.Save(entity).Error
}

func (r *repository[T]) Delete(entity *T) error {
	return r.db.Delete(entity).Error
}
