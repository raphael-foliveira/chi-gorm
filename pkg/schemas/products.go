package schemas

type CreateProduct struct {
	Name  string  `json:"name" faker:"name"`
	Price float64 `json:"price" faker:"amount"`
}

type Product struct {
	ID    uint    `json:"id"`
	Name  string  `json:"name" faker:"name"`
	Price float64 `json:"price" faker:"amount"`
}
