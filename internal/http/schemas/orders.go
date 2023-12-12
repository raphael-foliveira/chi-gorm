package schemas

import "github.com/raphael-foliveira/chi-gorm/internal/entities"

type CreateOrder struct {
	ClientID  int64 `json:"client_id" faker:"-"`
	ProductID int64 `json:"product_id" faker:"-"`
	Quantity  int   `json:"quantity"`
}

func (co *CreateOrder) ToModel() *entities.Order {
	return &entities.Order{
		ClientID:  co.ClientID,
		ProductID: co.ProductID,
		Quantity:  co.Quantity,
	}
}

type UpdateOrder struct {
	Quantity int `json:"quantity"`
}

type Order struct {
	ID       int64    `json:"id" faker:"-"`
	Quantity int      `json:"quantity" faker:"-"`
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
