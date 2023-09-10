package schemas

import "github.com/raphael-foliveira/chi-gorm/pkg/models"

type CreateClient struct {
	Name  string `json:"name" faker:"name"`
	Email string `json:"email" faker:"email"`
}

type UpdateClient CreateClient

type Client struct {
	ID    uint   `json:"id" faker:"-"`
	Name  string `json:"name" faker:"name"`
	Email string `json:"email" faker:"email"`
}

func NewClient(clientModel models.Client) *Client {
	return &Client{
		ID:    clientModel.ID,
		Name:  clientModel.Name,
		Email: clientModel.Email,
	}
}

func NewClients(clients []models.Client) []*Client {
	var c []*Client
	for _, client := range clients {
		c = append(c, NewClient(client))
	}
	return c
}

type ClientOrder struct {
	ID       uint     `json:"id" faker:"-"`
	Quantity uint     `json:"quantity" faker:"-"`
	Product  *Product `json:"product" faker:"-"`
}

func NewClientOrder(orderModel models.Order, productModel models.Product) *ClientOrder {
	return &ClientOrder{
		ID:       orderModel.ID,
		Quantity: orderModel.Quantity,
		Product:  NewProduct(productModel),
	}
}

type ClientDetail struct {
	ID     uint           `json:"id" faker:"-"`
	Name   string         `json:"name" faker:"name"`
	Email  string         `json:"email" faker:"email"`
	Orders []*ClientOrder `json:"orders" faker:"-"`
}

func NewClientDetail(clientModel models.Client, orders []models.Order) *ClientDetail {
	c := new(ClientDetail)
	c.ID = clientModel.ID
	c.Name = clientModel.Name
	c.Email = clientModel.Email
	for _, order := range orders {
		c.Orders = append(c.Orders, NewClientOrder(order, order.Product))
	}
	return c
}
