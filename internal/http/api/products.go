package api

import (
	"net/http"

	"github.com/raphael-foliveira/chi-gorm/internal/http/schemas"
	"github.com/raphael-foliveira/chi-gorm/internal/ports"
)

type ProductsController struct {
	productsRepo ports.ProductsRepository
}

func NewProductsController(productsRepo ports.ProductsRepository) *ProductsController {
	return &ProductsController{
		productsRepo: productsRepo,
	}
}

func (p *ProductsController) Create(ctx *Context) error {
	var body schemas.CreateProduct
	err := ctx.ParseBody(&body)
	if err != nil {
		return err
	}
	newProduct := body.ToModel()
	if err := p.productsRepo.Create(newProduct); err != nil {
		return err
	}
	return ctx.JSON(http.StatusCreated, schemas.NewProduct(newProduct))
}

func (p *ProductsController) Update(ctx *Context) error {
	id, err := ctx.GetUintPathParam("id")
	if err != nil {
		return err
	}
	var body schemas.UpdateProduct
	if err := ctx.ParseBody(&body); err != nil {
		return err
	}

	order := body.ToModel()
	order.ID = id

	if err := p.productsRepo.Update(order); err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, schemas.NewProduct(order))
}

func (p *ProductsController) Delete(ctx *Context) error {
	id, err := ctx.GetUintPathParam("id")
	if err != nil {
		return err
	}
	err = p.productsRepo.Delete(id)
	if err != nil {
		return err
	}
	return ctx.SendStatus(http.StatusNoContent)
}

func (p *ProductsController) List(ctx *Context) error {
	products, err := p.productsRepo.List()
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, schemas.NewProducts(products))
}

func (p *ProductsController) Get(ctx *Context) error {
	id, err := ctx.GetUintPathParam("id")
	if err != nil {
		return err
	}
	product, err := p.productsRepo.Get(id)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, schemas.NewProduct(product))
}
