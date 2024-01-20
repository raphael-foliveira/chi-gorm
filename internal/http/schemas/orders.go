package schemas

import (
	"github.com/raphael-foliveira/chi-gorm/internal/entities"
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

func (co *CreateOrder) Validate() (err error) {
	if co.Quantity <= 0 {
		err = addFieldError(err, "quantity", "quantity must be greater than zero")
	}
	return NewValidationError(err)
}

type UpdateOrder struct {
	Quantity uint `json:"quantity"`
}

func (uo *UpdateOrder) Validate() (err error) {
	if uo.Quantity <= 0 {
		err = addFieldError(err, "quantity", "quantity must be greater than zero")
	}
	return NewValidationError(err)
}

type Order struct {
	ID       uint     `json:"id" faker:"-"`
	Quantity uint     `json:"quantity" faker:"-"`
	Client   *Client  `json:"client" faker:"-"`
	Product  *Product `json:"product" faker:"-"`
}

func NewOrder(orderModel *entities.Order) *Order {
	return &Order{
		ID:       orderModel.ID,
		Quantity: orderModel.Quantity,
		Client:   NewClient(&orderModel.Client),
		Product:  NewProduct(&orderModel.Product),
	}
}

func NewOrders(orders []entities.Order) []Order {
	o := []Order{}
	for _, order := range orders {
		o = append(o, *NewOrder(&order))
	}
	return o
}
