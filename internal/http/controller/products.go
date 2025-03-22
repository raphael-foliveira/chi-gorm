package controller

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/raphael-foliveira/chi-gorm/internal/http/schemas"
	"github.com/raphael-foliveira/chi-gorm/internal/ports"
)

type Products struct {
	productsRepo ports.ProductsRepository
}

func NewProducts(productsRepo ports.ProductsRepository) *Products {
	return &Products{
		productsRepo: productsRepo,
	}
}

func (c *Products) Mount(mux chi.Router) {
	router := NewRouter()
	router.Get("/", c.List)
	router.Post("/", c.Create)
	router.Get("/{id}", c.Get)
	router.Delete("/{id}", c.Delete)
	router.Put("/{id}", c.Update)
	mux.Mount("/products", router)
}

func (p *Products) Create(ctx *Context) error {
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

func (p *Products) Update(ctx *Context) error {
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

func (p *Products) Delete(ctx *Context) error {
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

func (p *Products) List(ctx *Context) error {
	products, err := p.productsRepo.List()
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, schemas.NewProducts(products))
}

func (p *Products) Get(ctx *Context) error {
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
