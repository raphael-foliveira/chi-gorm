package client

import (
	"github.com/raphael-foliveira/chi-gorm/internal/db"
)

type iRepository interface {
	List() ([]Client, error)
	Get(id uint64) (Client, error)
	Create(c *Client) error
	Update(c *Client) error
	Delete(c *Client) error
}

type Repository struct {
	db *db.DB
}

func NewRepository(db *db.DB) *Repository {
	return &Repository{db}
}

func (r *Repository) List() ([]Client, error) {
	c := []Client{}
	return c, r.db.Find(&c).Error
}

func (r *Repository) Get(id uint64) (Client, error) {
	c := Client{}
	return c, r.db.First(&c, id).Error
}

func (r *Repository) Create(c *Client) error {
	return r.db.Create(c).Error
}

func (r *Repository) Update(c *Client) error {
	return r.db.Save(c).Error
}

func (r *Repository) Delete(c *Client) error {
	return r.db.Delete(c).Error
}
