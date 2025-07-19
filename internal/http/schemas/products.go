package schemas

import (
	"github.com/raphael-foliveira/chi-gorm/internal/entities"
	"github.com/raphael-foliveira/chi-gorm/internal/validation"
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

func (cp *CreateProduct) Validate() map[string][]string {
	return validation.Validate(func(v *validation.Validator) {
		v.Check("name", cp.Name != "", "name is required")
		v.Check("price", cp.Price > 0, "price must be greater than 0")
	})
}

type UpdateProduct struct {
	CreateProduct
}

type Product struct {
	Name  string  `json:"name" faker:"name"`
	ID    uint    `json:"id"`
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
	p := []Product{}
	for _, product := range e {
		p = append(p, *NewProduct(&product))
	}
	return p
}
