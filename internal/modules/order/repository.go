package order

import (
	"github.com/raphael-foliveira/chi-gorm/internal/db"
	"github.com/raphael-foliveira/chi-gorm/internal/modules/client"
	"github.com/raphael-foliveira/chi-gorm/internal/modules/product"
)

type Repository struct {
	db *db.DB
}

func NewRepository(db *db.DB) *Repository {
	return &Repository{db}
}

func (r *Repository) List() ([]Order, error) {
	o := []Order{}
	return o, r.db.Preload("Client").Preload("Product").Find(&o).Error
}

func (r *Repository) Get(id uint64) (Order, error) {
	o := Order{}
	return o, r.db.Preload("Client").Preload("Product").First(&o, id).Error
}

func (r *Repository) Create(c *Order) error {
	cli := client.Client{}
	prod := product.Product{}
	err := r.db.First(&cli, c.ClientID).Error
	if err != nil {
		return err
	}
	err = r.db.First(&prod, c.ProductID).Error
	if err != nil {
		return err
	}
	c.Client = cli
	c.Product = prod
	return r.db.Create(c).Error
}

func (r *Repository) Update(c *Order) error {
	return r.db.Save(c).Error
}

func (r *Repository) Delete(c *Order) error {
	return r.db.Delete(c).Error
}
