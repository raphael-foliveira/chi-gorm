package controller

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/raphael-foliveira/chi-gorm/internal/http/schemas"
)

type Products struct{}

func NewProducts() *Products {
	return &Products{}
}

func (p *Products) Mount() {
	router := chi.NewRouter()
	router.Get("/", useHandler(p.List))
	router.Post("/", useHandler(p.Create))
	router.Get("/{id}", useHandler(p.Get))
	router.Delete("/{id}", useHandler(p.Delete))
	router.Put("/{id}", useHandler(p.Update))

	app.Mount("/products", router)
}

func (p *Products) Create(ctx *Context) error {
	body := &schemas.CreateProduct{}
	err := ctx.ParseBody(body)
	if err != nil {
		return err
	}
	newOrder, err := productsService.Create(body)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusCreated, schemas.NewProduct(newOrder))
}

func (p *Products) Update(ctx *Context) error {
	id, err := ctx.GetUintPathParam("id")
	if err != nil {
		return err
	}
	body := &schemas.UpdateProduct{}
	err = ctx.ParseBody(body)
	if err != nil {
		return err
	}
	updatedOrder, err := productsService.Update(id, body)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, schemas.NewProduct(updatedOrder))
}

func (p *Products) Delete(ctx *Context) error {
	id, err := ctx.GetUintPathParam("id")
	if err != nil {
		return err
	}
	err = productsService.Delete(id)
	if err != nil {
		return err
	}
	return ctx.SendStatus(http.StatusNoContent)
}

func (p *Products) List(ctx *Context) error {
	products, err := productsService.List()
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
	product, err := productsService.Get(id)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, schemas.NewProduct(product))
}
