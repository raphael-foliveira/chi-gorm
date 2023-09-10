package schemas

type CreateOrder struct {
	ClientID  uint `json:"client_id" faker:"-"`
	ProductID uint `json:"product_id" faker:"-"`
	Quantity  uint `json:"quantity" faker:"-"`
}

type Order struct {
	ID        uint `json:"id" faker:"-"`
	ClientID  uint `json:"client_id" faker:"-"`
	ProductID uint `json:"product_id" faker:"-"`
	Quantity  uint `json:"quantity" faker:"-"`
}

type OrderDetail struct {
	ID        uint    `json:"id" faker:"-"`
	ClientID  uint    `json:"client_id" faker:"-"`
	ProductID uint    `json:"product_id" faker:"-"`
	Quantity  uint    `json:"quantity" faker:"-"`
	Client    Client  `json:"client" faker:"-"`
	Product   Product `json:"product" faker:"-"`
}
