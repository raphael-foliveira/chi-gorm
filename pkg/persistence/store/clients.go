package store

import (
	"github.com/raphael-foliveira/chi-gorm/pkg/interfaces"
	"github.com/raphael-foliveira/chi-gorm/pkg/models"
	"github.com/raphael-foliveira/chi-gorm/pkg/persistence/db"
)

type Clients interface {
	interfaces.Store[models.Client]
}

type clients struct{}

func NewClients() Clients {
	db.Db.AutoMigrate(&models.Client{})
	return &clients{}
}

func (r *clients) List() ([]models.Client, error) {
	clients := []models.Client{}
	return clients, db.Db.Find(&clients).Error
}

func (r *clients) Get(id int64) (*models.Client, error) {
	client := models.Client{}
	return &client, db.Db.First(&client, id).Error
}

func (r *clients) Create(client *models.Client) error {
	return db.Db.Create(client).Error
}

func (r *clients) Update(client *models.Client) error {
	return db.Db.Save(client).Error
}

func (r *clients) Delete(client *models.Client) error {
	return db.Db.Delete(client).Error
}
