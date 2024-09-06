package repository

import (
	"log"

	"gorm.io/gorm"
)

type Repository[T any] struct{}

func New[T any]() *Repository[T] {
	return &Repository[T]{}
}

func (r *Repository[T]) List(conds ...any) ([]T, error) {
	entities := []T{}
	log.Printf("db: %+v\n", db)
	return entities, db.Find(&entities, conds...).Error
}

func (r *Repository[T]) Get(id uint) (*T, error) {
	entity := new(T)
	return entity, db.Model(new(T)).First(entity, id).Error
}

func (r *Repository[T]) Create(entity *T) error {
	return db.Create(entity).Error
}

func (r *Repository[T]) Update(entity *T) error {
	return db.Save(entity).Error
}

func (r *Repository[T]) Delete(entity *T) error {
	return db.Delete(entity).Error
}

var db *gorm.DB

func Initialize(gormDb *gorm.DB) {
	db = gormDb
}
