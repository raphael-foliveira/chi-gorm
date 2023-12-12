package repository

import (
	"github.com/raphael-foliveira/chi-gorm/internal/entities"
	"github.com/raphael-foliveira/chi-gorm/internal/interfaces"
	"gorm.io/gorm"
)

type Clients interface {
	interfaces.Repository[entities.Client]
}

type clients struct {
	db *gorm.DB
}

func NewClients(db *gorm.DB) Clients {
	db.AutoMigrate(&entities.Client{})
	return &clients{db}
}

func (r *clients) List() ([]entities.Client, error) {
	clients := []entities.Client{}
	return clients, r.db.Find(&clients).Error
}

func (r *clients) Get(id int64) (*entities.Client, error) {
	client := entities.Client{}
	return &client, r.db.Model(&entities.Client{}).Preload("Orders.Product").First(&client, id).Error
}

func (r *clients) Create(client *entities.Client) error {
	return r.db.Create(client).Error
}

func (r *clients) Update(client *entities.Client) error {
	return r.db.Save(client).Error
}

func (r *clients) Delete(client *entities.Client) error {
	return r.db.Delete(client).Error
}
