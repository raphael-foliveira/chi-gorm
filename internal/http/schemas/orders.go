package schemas

import (
	"fmt"

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
	CreateOrder
}

func (uo *UpdateOrder) Validate() error {
	return validate.Rules(func() error {
		if int(uo.Quantity) < 1 {
			return fmt.Errorf("%v: must be greater than %v", "quantity", 1)
		}
		return nil
	}())
}

func (uo *UpdateOrder) ToModel() *entities.Order {
	return &entities.Order{
		Quantity:  uo.Quantity,
		ClientID:  uo.ClientID,
		ProductID: uo.ProductID,
	}
}

type Order struct {
	Client   *Client  `json:"client"`
	Product  *Product `json:"product"`
	ID       uint     `json:"id"`
	Quantity uint     `json:"quantity"`
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
