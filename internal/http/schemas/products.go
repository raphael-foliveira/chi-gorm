package schemas

import (
	"fmt"

	"github.com/raphael-foliveira/chi-gorm/internal/entities"
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

func (cp *CreateProduct) Validate() error {
	err := NewValidationError()
	if cp.Name == "" {
		err.Add(errProductNameInvalid)
	}
	if cp.Price <= 0 {
		err.Add(errPriceInvalid)
	}
	return err.ReturnIfError()
}

type UpdateProduct struct {
	CreateProduct
}

type Product struct {
	ID    uint    `json:"id"`
	Name  string  `json:"name" faker:"name"`
	Price float64 `json:"price" faker:"amount"`
}

func NewProduct(productModel *entities.Product) *Product {
	return &Product{
		ID:    productModel.ID,
		Name:  productModel.Name,
		Price: productModel.Price,
	}
}

func NewProducts(products []entities.Product) []Product {
	p := []Product{}
	for _, product := range products {
		p = append(p, *NewProduct(&product))
	}
	return p
}

var errProductNameInvalid = fmt.Errorf("%w: name is required", ErrValidation)
var errPriceInvalid = fmt.Errorf("%w: price must be greater than zero", ErrValidation)
