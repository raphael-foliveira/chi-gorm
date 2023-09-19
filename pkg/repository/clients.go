package repository

import (
	"github.com/raphael-foliveira/chi-gorm/pkg/database"
	"github.com/raphael-foliveira/chi-gorm/pkg/interfaces"
	"github.com/raphael-foliveira/chi-gorm/pkg/models"
	"gorm.io/gorm"
)

type Clients interface {
	interfaces.Repository[models.Client]
}

type clients struct {
	db *gorm.DB
}

func NewClients() Clients {
	db := database.GetDb()
	db.AutoMigrate(&models.Client{})
	return &clients{db}
}

func (r *clients) List() ([]models.Client, error) {
	clients := []models.Client{}
	return clients, r.db.Find(&clients).Error
}

func (r *clients) Get(id int64) (*models.Client, error) {
	client := models.Client{}
	return &client, r.db.First(&client, id).Error
}

func (r *clients) Create(client *models.Client) error {
	return r.db.Create(client).Error
}

func (r *clients) Update(client *models.Client) error {
	return r.db.Save(client).Error
}

func (r *clients) Delete(client *models.Client) error {
	return r.db.Delete(client).Error
}
