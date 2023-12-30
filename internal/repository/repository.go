package repository

import "github.com/raphael-foliveira/chi-gorm/internal/database"

type Repository[T interface{}] interface {
	List() ([]T, error)
	Get(uint) (*T, error)
	Create(*T) error
	Update(*T) error
	Delete(*T) error
}

type repository[T interface{}] struct{}

func (r *repository[T]) List() ([]T, error) {
	entities := []T{}
	return entities, database.Db.Find(&entities).Error
}

func (r *repository[T]) Get(id uint) (*T, error) {
	entity := new(T)
	return entity, database.Db.Model(new(T)).First(entity, id).Error
}

func (r *repository[T]) Create(entity *T) error {
	return database.Db.Create(entity).Error
}

func (r *repository[T]) Update(entity *T) error {
	return database.Db.Save(entity).Error
}

func (r *repository[T]) Delete(entity *T) error {
	return database.Db.Delete(entity).Error
}
