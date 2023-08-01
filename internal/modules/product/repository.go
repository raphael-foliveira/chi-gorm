package product

import (
	"github.com/raphael-foliveira/chi-gorm/internal/db"
)

type iRepository interface {
	List() ([]Product, error)
	Get(id uint64) (Product, error)
	Create(c *Product) error
	Update(c *Product) error
	Delete(c *Product) error
}

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
	c := Product{}
	return c, r.db.First(&c, id).Error
}

func (r *Repository) Create(c *Product) error {
	return r.db.Create(c).Error
}

func (r *Repository) Update(c *Product) error {
	return r.db.Save(c).Error
}

func (r *Repository) Delete(c *Product) error {
	return r.db.Delete(c).Error
}
