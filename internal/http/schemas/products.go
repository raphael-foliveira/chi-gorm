package schemas

import (
	"github.com/raphael-foliveira/chi-gorm/internal/entities"
	"github.com/raphael-foliveira/chi-gorm/internal/validate"
)

type CreateProduct struct {
	Name  string  `json:"name" faker:"name"`
	Price float64 `json:"price" faker:"amount"`
}

func (cp *CreateProduct) ToModel() *entities.Product {
	return &entities.Product{
		Name:  cp.Name,
		Price: cp.Price,
	}
}

func (cp *CreateProduct) Validate() (err error) {
	return validate.Rules(validate.Required("name", cp.Name), validate.Min("price", int(cp.Price), 0))
}

type UpdateProduct struct {
	CreateProduct
}

type Product struct {
	ID    uint    `json:"id"`
	Name  string  `json:"name" faker:"name"`
	Price float64 `json:"price" faker:"amount"`
}

func NewProduct(e *entities.Product) *Product {
	return &Product{
		ID:    e.ID,
		Name:  e.Name,
		Price: e.Price,
	}
}

func NewProducts(e []entities.Product) []Product {
	ps := []Product{}
	for _, product := range e {
		ps = append(ps, *NewProduct(&product))
	}
	return ps
}
