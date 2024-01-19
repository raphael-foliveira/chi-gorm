package schemas

import (
	"errors"
	"strings"

	"github.com/raphael-foliveira/chi-gorm/internal/entities"
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
	if cc.Name == "" {
		err = errors.Join(err, errClientNameRequired)
	}
	if cc.Email == "" {
		err = errors.Join(err, errEmailRequired)
	}
	if !strings.Contains(cc.Email, "@") {
		err = errors.Join(err, errEmailInvalid)
	}
	return NewValidationErrors(err)
}

type UpdateClient struct {
	CreateClient
}

type Client struct {
	ID    uint   `json:"id" faker:"-"`
	Name  string `json:"name" faker:"name"`
	Email string `json:"email" faker:"email"`
}

func NewClient(clientModel *entities.Client) *Client {
	return &Client{
		ID:    clientModel.ID,
		Name:  clientModel.Name,
		Email: clientModel.Email,
	}
}

func NewClients(clients []entities.Client) []Client {
	c := []Client{}
	for _, client := range clients {
		c = append(c, *NewClient(&client))
	}
	return c
}

type ClientOrder struct {
	ID       uint     `json:"id" faker:"-"`
	Quantity uint     `json:"quantity" faker:"-"`
	Product  *Product `json:"product" faker:"-"`
}

func NewClientOrder(orderModel *entities.Order) *ClientOrder {
	return &ClientOrder{
		ID:       orderModel.ID,
		Quantity: orderModel.Quantity,
		Product:  NewProduct(&orderModel.Product),
	}
}

func NewClientOrders(orders []entities.Order) []ClientOrder {
	o := []ClientOrder{}
	for _, order := range orders {
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

func NewClientDetail(clientModel *entities.Client) *ClientDetail {
	c := &ClientDetail{}
	c.ID = clientModel.ID
	c.Name = clientModel.Name
	c.Email = clientModel.Email
	c.Orders = NewClientOrders(clientModel.Orders)
	return c
}

var errClientNameRequired = errors.New("client: client name is required")
var errEmailRequired = errors.New("email: email is required")
var errEmailInvalid = errors.New("email: invalid email")
