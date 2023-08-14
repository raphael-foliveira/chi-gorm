package order

import (
	"github.com/raphael-foliveira/chi-gorm/internal/db"
	"github.com/raphael-foliveira/chi-gorm/internal/modules/client"
	"github.com/raphael-foliveira/chi-gorm/internal/modules/product"
)

type iRepository interface {
	List() ([]Order, error)
	Get(id uint64) (Order, error)
	Create(c *Order) error
	Update(c *Order) error
	Delete(c *Order) error
}

type Repository struct {
	db *db.DB
}

func NewRepository(db *db.DB) *Repository {
	return &Repository{db}
}

func (r *Repository) List() ([]Order, error) {
	c := []Order{}
	return c, r.db.Preload("Client").Preload("Product").Find(&c).Error
}

func (r *Repository) Get(id uint64) (Order, error) {
	c := Order{}
	return c, r.db.Preload("Client").Preload("Product").First(&c, id).Error
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
