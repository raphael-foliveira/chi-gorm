package client

import (
	"github.com/raphael-foliveira/chi-gorm/internal/db"
	"github.com/raphael-foliveira/chi-gorm/internal/modules/order"
)

type Repository struct {
	db *db.DB
}

func NewRepository(db *db.DB) *Repository {
	return &Repository{db}
}

func (r *Repository) List() ([]Client, error) {
	clients := []Client{}
	err := r.db.Find(&clients).Error
	if err != nil {
		return nil, err
	}
	for i := range clients {
		orders, err := order.NewRepository(r.db).ListByClient(clients[i].ID)
		if err != nil {
			return nil, err
		}
		clients[i].Orders = orders
	}
	return clients, nil
}

func (r *Repository) Get(id uint64) (Client, error) {
	client := Client{}
	return client, r.db.First(&client, id).Error
}

func (r *Repository) Create(client *Client) error {
	return r.db.Create(client).Error
}

func (r *Repository) Update(client *Client) error {
	return r.db.Save(client).Error
}

func (r *Repository) Delete(client *Client) error {
	return r.db.Delete(client).Error
}
