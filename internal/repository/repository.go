package repository

import (
	"gorm.io/gorm"
)

type Repository[T interface{}] interface {
	List() ([]T, error)
	Get(id int64) (*T, error)
	Create(c *T) error
	Update(c *T) error
	Delete(c *T) error
}

type repository[T interface{}] struct {
	db *gorm.DB
}

func newRepository[T interface{}](db *gorm.DB) *repository[T] {
	return &repository[T]{db}

}

func (r *repository[T]) List() ([]T, error) {
	entities := []T{}
	return entities, r.db.Find(&entities).Error
}

func (r *repository[T]) Get(id int64) (*T, error) {
	entity := new(T)
	return entity, r.db.Model(new(T)).First(&entity, id).Error
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
