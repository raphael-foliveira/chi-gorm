package product

import (
	"github.com/raphael-foliveira/chi-gorm/internal/db"
)

type Repository struct {
	db *db.DB
}

func NewRepository(db *db.DB) *Repository {
	return &Repository{db}
}

func (r *Repository) List() ([]Product, error) {
	c := []Product{}
	return c, r.db.Find(&c).Error
}

func (r *Repository) Get(id uint64) (Product, error) {
	product := Product{}
	return product, r.db.First(&product, id).Error
}

func (r *Repository) Create(product *Product) error {
	return r.db.Create(product).Error
}

func (r *Repository) Update(product *Product) error {
	return r.db.Save(product).Error
}

func (r *Repository) Delete(product *Product) error {
	return r.db.Delete(product).Error
}
