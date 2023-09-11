package schemas

import "github.com/raphael-foliveira/chi-gorm/pkg/models"

type CreateOrder struct {
	ClientID  uint `json:"client_id" faker:"-"`
	ProductID uint `json:"product_id" faker:"-"`
	Quantity  int  `json:"quantity" faker:"-"`
}

type UpdateOrder struct {
	Quantity int `json:"quantity" faker:"-"`
}

type Order struct {
	ID        uint `json:"id" faker:"-"`
	ClientID  uint `json:"client_id" faker:"-"`
	ProductID uint `json:"product_id" faker:"-"`
	Quantity  int  `json:"quantity" faker:"-"`
}

func NewOrder(orderModel models.Order) *Order {
	return &Order{
		ID:        orderModel.ID,
		ClientID:  orderModel.ClientID,
		ProductID: orderModel.ProductID,
		Quantity:  orderModel.Quantity,
	}
}

type OrderDetail struct {
	ID        uint     `json:"id" faker:"-"`
	ClientID  uint     `json:"client_id" faker:"-"`
	ProductID uint     `json:"product_id" faker:"-"`
	Quantity  int      `json:"quantity" faker:"-"`
	Client    *Client  `json:"client" faker:"-"`
	Product   *Product `json:"product" faker:"-"`
}

func NewOrderDetail(orderModel models.Order, clientModel models.Client, productModel models.Product) *OrderDetail {
	return &OrderDetail{
		ID:        orderModel.ID,
		ClientID:  orderModel.ClientID,
		ProductID: orderModel.ProductID,
		Quantity:  orderModel.Quantity,
		Client:    NewClient(clientModel),
		Product:   NewProduct(productModel),
	}
}
