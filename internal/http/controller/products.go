package controller

import (
	"net/http"

	"github.com/raphael-foliveira/chi-gorm/internal/http/schemas"
	"github.com/raphael-foliveira/chi-gorm/internal/service"
)

type products struct{}

func NewProducts() *products {
	return &products{}
}

func (p *products) Create(ctx *Context) error {
	body := &schemas.CreateProduct{}
	err := ctx.ParseBody(body)
	if err != nil {
		return err
	}
	newOrder, err := service.Products.Create(body)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusCreated, schemas.NewProduct(newOrder))
}

func (p *products) Update(ctx *Context) error {
	id, err := ctx.GetUintPathParam("id")
	if err != nil {
		return err
	}
	body := &schemas.UpdateProduct{}
	err = ctx.ParseBody(body)
	if err != nil {
		return err
	}
	updatedOrder, err := service.Products.Update(id, body)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, schemas.NewProduct(updatedOrder))
}

func (p *products) Delete(ctx *Context) error {
	id, err := ctx.GetUintPathParam("id")
	if err != nil {
		return err
	}
	err = service.Products.Delete(id)
	if err != nil {
		return err
	}
	return ctx.SendStatus(http.StatusNoContent)
}

func (p *products) List(ctx *Context) error {
	products, err := service.Products.List()
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, schemas.NewProducts(products))
}

func (p *products) Get(ctx *Context) error {
	id, err := ctx.GetUintPathParam("id")
	if err != nil {
		return err
	}
	product, err := service.Products.Get(id)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, schemas.NewProduct(product))
}
