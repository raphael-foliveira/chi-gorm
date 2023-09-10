package schemas

type CreateClient struct {
	Name  string `json:"name" faker:"name"`
	Email string `json:"email" faker:"email"`
}

type Client struct {
	ID    uint   `json:"id" faker:"-"`
	Name  string `json:"name" faker:"name"`
	Email string `json:"email" faker:"email"`
}

type ClientDetail struct {
	ID     uint          `json:"id" faker:"-"`
	Name   string        `json:"name" faker:"name"`
	Email  string        `json:"email" faker:"email"`
	Orders []ClientOrder `json:"orders" faker:"-"`
}

type ClientOrder struct {
	ID       uint    `json:"id" faker:"-"`
	Quantity uint    `json:"quantity" faker:"-"`
	Product  Product `json:"product" faker:"-"`
}
