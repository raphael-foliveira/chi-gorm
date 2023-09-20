package schemas

import (
	"github.com/raphael-foliveira/chi-gorm/pkg/models"
)

type CreateClient struct {
	Name  string `json:"name" faker:"name"`
	Email string `json:"email" faker:"email"`
}

func (cc *CreateClient) ToModel() models.Client {
	return models.Client{
		Name:  cc.Name,
		Email: cc.Email,
	}
}

type UpdateClient struct {
	CreateClient
}

type Client struct {
	ID    int64  `json:"id" faker:"-"`
	Name  string `json:"name" faker:"name"`
	Email string `json:"email" faker:"email"`
}

func NewClient(clientModel *models.Client) *Client {
	return &Client{
		ID:    clientModel.ID,
		Name:  clientModel.Name,
		Email: clientModel.Email,
	}
}

func NewClients(clients []models.Client) []*Client {
	c := []*Client{}
	for _, client := range clients {
		c = append(c, NewClient(&client))
	}
	return c
}

type ClientOrder struct {
	ID       int64    `json:"id" faker:"-"`
	Quantity int      `json:"quantity" faker:"-"`
	Product  *Product `json:"product" faker:"-"`
}

func NewClientOrder(orderModel *models.Order) *ClientOrder {
	return &ClientOrder{
		ID:       orderModel.ID,
		Quantity: orderModel.Quantity,
		Product:  NewProduct(&orderModel.Product),
	}
}

func NewClientOrders(orders []models.Order) []*ClientOrder {
	o := []*ClientOrder{}
	for _, order := range orders {
		o = append(o, NewClientOrder(&order))
	}
	return o
}

type ClientDetail struct {
	ID     int64          `json:"id" faker:"-"`
	Name   string         `json:"name" faker:"name"`
	Email  string         `json:"email" faker:"email"`
	Orders []*ClientOrder `json:"orders" faker:"-"`
}

func NewClientDetail(clientModel *models.Client) *ClientDetail {
	c := &ClientDetail{}
	c.ID = clientModel.ID
	c.Name = clientModel.Name
	c.Email = clientModel.Email
	c.Orders = NewClientOrders(clientModel.Orders)
	return c
}
