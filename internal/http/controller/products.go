package controller

import (
	"net/http"

	"github.com/raphael-foliveira/chi-gorm/internal/http/schemas"
	"github.com/raphael-foliveira/chi-gorm/internal/service"
)

type Products struct {
	service *service.Products
}

func NewProducts(service *service.Products) *Products {
	return &Products{service}
}

func (p *Products) Create(ctx *Context) error {
	body := &schemas.CreateProduct{}
	err := ctx.ParseBody(body)
	if err != nil {
		return err
	}
	newOrder, err := p.service.Create(body)
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
	updatedOrder, err := p.service.Update(id, body)
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
	err = p.service.Delete(id)
	if err != nil {
		return err
	}
	return ctx.SendStatus(http.StatusNoContent)
}

func (p *Products) List(ctx *Context) error {
	products, err := p.service.List()
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
	product, err := p.service.Get(id)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, schemas.NewProduct(product))
}
