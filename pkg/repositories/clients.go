package repositories

import (
	"github.com/raphael-foliveira/chi-gorm/pkg/db"
	"github.com/raphael-foliveira/chi-gorm/pkg/interfaces"
	"github.com/raphael-foliveira/chi-gorm/pkg/models"
)

type Clients interface {
	interfaces.Repository[models.Client]
}

type clients struct {
	db *db.DB
}

func NewClient(db *db.DB) *clients {
	return &clients{db}
}

func (r *clients) List() ([]models.Client, error) {
	clients := []models.Client{}
	err := r.db.Find(&clients).Error
	if err != nil {
		return nil, err
	}
	return clients, nil
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
