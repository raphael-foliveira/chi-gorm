package schemas

import (
	"github.com/raphael-foliveira/chi-gorm/internal/entities"
	"github.com/raphael-foliveira/chi-gorm/internal/validate"
)

type CreateClient struct {
	Name  string `json:"name" faker:"name"`
	Email string `json:"email" faker:"email"`
}

func (cc *CreateClient) ToModel() *entities.Client {
	return &entities.Client{
		Name:  cc.Name,
		Email: cc.Email,
	}
}

func (cc *CreateClient) Validate() error {
	return validate.Rules(
		validate.Required("name", cc.Name),
		validate.Required("email", cc.Email),
		validate.Email("email", cc.Email),
	)
}

type UpdateClient struct {
	CreateClient
}

type Client struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	ID    uint   `json:"id"`
}

func NewClient(e *entities.Client) *Client {
	return &Client{
		ID:    e.ID,
		Name:  e.Name,
		Email: e.Email,
	}
}

func NewClients(e []entities.Client) []Client {
	c := []Client{}
	for _, client := range e {
		c = append(c, *NewClient(&client))
	}
	return c
}

type ClientOrder struct {
	Product  *Product `json:"product"`
	ID       uint     `json:"id"`
	Quantity uint     `json:"quantity"`
}

func NewClientOrder(e *entities.Order) *ClientOrder {
	return &ClientOrder{
		ID:       e.ID,
		Quantity: e.Quantity,
		Product:  NewProduct(&e.Product),
	}
}

func NewClientOrders(e []entities.Order) []ClientOrder {
	o := []ClientOrder{}
	for _, order := range e {
		o = append(o, *NewClientOrder(&order))
	}
	return o
}

type ClientDetail struct {
	Name   string        `json:"name" faker:"name"`
	Email  string        `json:"email" faker:"email"`
	Orders []ClientOrder `json:"orders"`
	ID     uint          `json:"id"`
}

func NewClientDetail(e *entities.Client) *ClientDetail {
	return &ClientDetail{
		ID:     e.ID,
		Name:   e.Name,
		Email:  e.Email,
		Orders: NewClientOrders(e.Orders),
	}
}
