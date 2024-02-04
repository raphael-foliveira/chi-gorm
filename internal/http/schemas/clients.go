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

func (cc *CreateClient) Validate() (err error) {
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
	ID    uint   `json:"id" faker:"-"`
	Name  string `json:"name" faker:"name"`
	Email string `json:"email" faker:"email"`
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
	ID       uint     `json:"id" faker:"-"`
	Quantity uint     `json:"quantity" faker:"-"`
	Product  *Product `json:"product" faker:"-"`
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
	ID     uint          `json:"id" faker:"-"`
	Name   string        `json:"name" faker:"name"`
	Email  string        `json:"email" faker:"email"`
	Orders []ClientOrder `json:"orders" faker:"-"`
}

func NewClientDetail(e *entities.Client) *ClientDetail {
	return &ClientDetail{
		ID:     e.ID,
		Name:   e.Name,
		Email:  e.Email,
		Orders: NewClientOrders(e.Orders),
	}
}
