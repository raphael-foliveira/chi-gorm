package schemas

import "github.com/raphael-foliveira/chi-gorm/pkg/models"

type CreateProduct struct {
	Name  string  `json:"name" faker:"name"`
	Price float64 `json:"price" faker:"amount"`
}

type UpdateProduct struct {
	Name  string  `json:"name" faker:"name"`
	Price float64 `json:"price" faker:"amount"`
}

type Product struct {
	ID    int64   `json:"id"`
	Name  string  `json:"name" faker:"name"`
	Price float64 `json:"price" faker:"amount"`
}

func NewProduct(productModel models.Product) Product {
	return Product{
		ID:    productModel.ID,
		Name:  productModel.Name,
		Price: productModel.Price,
	}
}

func NewProducts(products []models.Product) []Product {
	var p []Product
	for _, product := range products {
		p = append(p, NewProduct(product))
	}
	return p
}
