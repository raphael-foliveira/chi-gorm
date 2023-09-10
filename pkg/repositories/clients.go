package repositories

import (
	"github.com/raphael-foliveira/chi-gorm/pkg/db"
	"github.com/raphael-foliveira/chi-gorm/pkg/models"
)

type Client struct {
	db *db.DB
}

func NewClient(db *db.DB) *Client {
	return &Client{db}
}

func (r *Client) List() ([]models.Client, error) {
	clients := []models.Client{}
	err := r.db.Find(&clients).Error
	if err != nil {
		return nil, err
	}
	return clients, nil
}

func (r *Client) Get(id int64) (models.Client, error) {
	client := models.Client{}
	return client, r.db.First(&client, id).Error
}

func (r *Client) Create(client *models.Client) error {
	return r.db.Create(client).Error
}

func (r *Client) Update(client *models.Client) error {
	return r.db.Save(client).Error
}

func (r *Client) Delete(client *models.Client) error {
	return r.db.Delete(client).Error
}
