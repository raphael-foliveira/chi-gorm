package schemas

import "github.com/raphael-foliveira/chi-gorm/pkg/models"

type CreateOrder struct {
	ClientID  int64 `json:"client_id" faker:"-"`
	ProductID int64 `json:"product_id" faker:"-"`
	Quantity  int   `json:"quantity"`
}

type UpdateOrder struct {
	Quantity int `json:"quantity"`
}

type Order struct {
	ID        int64 `json:"id" faker:"-"`
	ClientID  int64 `json:"client_id" faker:"-"`
	ProductID int64 `json:"product_id" faker:"-"`
	Quantity  int   `json:"quantity"`
}

func NewOrder(orderModel models.Order) Order {
	return Order{
		ID:        orderModel.ID,
		ClientID:  orderModel.ClientID,
		ProductID: orderModel.ProductID,
		Quantity:  orderModel.Quantity,
	}
}

func NewOrders(orders []models.Order) []Order {
	var o []Order
	for _, order := range orders {
		o = append(o, NewOrder(order))
	}
	return o
}

type OrderDetail struct {
	ID        int64   `json:"id" faker:"-"`
	ClientID  int64   `json:"client_id" faker:"-"`
	ProductID int64   `json:"product_id" faker:"-"`
	Quantity  int     `json:"quantity" faker:"-"`
	Client    Client  `json:"client" faker:"-"`
	Product   Product `json:"product" faker:"-"`
}

func NewOrderDetail(orderModel models.Order, clientModel models.Client, productModel models.Product) OrderDetail {
	return OrderDetail{
		ID:        orderModel.ID,
		ClientID:  orderModel.ClientID,
		ProductID: orderModel.ProductID,
		Quantity:  orderModel.Quantity,
		Client:    NewClient(clientModel),
		Product:   NewProduct(productModel),
	}
}
