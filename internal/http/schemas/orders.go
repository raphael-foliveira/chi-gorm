package schemas

import (
	"github.com/raphael-foliveira/chi-gorm/internal/entities"
	"github.com/raphael-foliveira/chi-gorm/internal/validate"
)

type CreateOrder struct {
	ClientID  uint `json:"client_id" faker:"-"`
	ProductID uint `json:"product_id" faker:"-"`
	Quantity  uint `json:"quantity"`
}

func (co *CreateOrder) ToModel() *entities.Order {
	return &entities.Order{
		ClientID:  co.ClientID,
		ProductID: co.ProductID,
		Quantity:  co.Quantity,
	}
}

func (co *CreateOrder) Validate() error {
	return validate.Rules(validate.Min("quantity", int(co.Quantity), 1))
}

type UpdateOrder struct {
	Quantity uint `json:"quantity"`
}

func (uo *UpdateOrder) Validate() error {
	return validate.Rules(validate.Min("quantity", int(uo.Quantity), 1))
}

type Order struct {
	Client   *Client  `json:"client" faker:"-"`
	Product  *Product `json:"product" faker:"-"`
	ID       uint     `json:"id" faker:"-"`
	Quantity uint     `json:"quantity" faker:"-"`
}

func NewOrder(e *entities.Order) *Order {
	return &Order{
		ID:       e.ID,
		Quantity: e.Quantity,
		Client:   NewClient(&e.Client),
		Product:  NewProduct(&e.Product),
	}
}

func NewOrders(e []entities.Order) []Order {
	os := []Order{}
	for _, order := range e {
		os = append(os, *NewOrder(&order))
	}
	return os
}
