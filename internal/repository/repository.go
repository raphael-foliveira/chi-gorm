package repository

import "gorm.io/gorm"

var db *gorm.DB

func SetDB(database *gorm.DB) {
	db = database
}

type repository[T any] struct{}

func newRepository[T any]() *repository[T] {
	return &repository[T]{}
}

func (r *repository[T]) List(conds ...any) ([]T, error) {
	entities := []T{}
	return entities, db.Find(&entities, conds...).Error
}

func (r *repository[T]) Get(id uint) (*T, error) {
	entity := new(T)
	return entity, db.Model(new(T)).First(entity, id).Error
}

func (r *repository[T]) Create(entity *T) error {
	return db.Create(entity).Error
}

func (r *repository[T]) Update(entity *T) error {
	return db.Save(entity).Error
}

func (r *repository[T]) Delete(entity *T) error {
	return db.Delete(entity).Error
}
