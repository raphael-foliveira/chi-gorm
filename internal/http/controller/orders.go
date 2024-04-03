package controller

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/raphael-foliveira/chi-gorm/internal/http/schemas"
	"github.com/raphael-foliveira/chi-gorm/internal/service"
)

type Orders struct {
	Service *service.Orders
	*router
}

func NewOrders(service *service.Orders) *Orders {
	router := &router{chi.NewRouter()}
	c := &Orders{service, router}
	router.Get("/", c.List)
	router.Post("/", c.Create)
	router.Get("/{id}", c.Get)
	router.Delete("/{id}", c.Delete)
	router.Put("/{id}", c.Update)
	return c
}

func (o *Orders) Create(ctx *Context) error {
	body := &schemas.CreateOrder{}
	err := ctx.ParseBody(body)
	if err != nil {
		return err
	}
	newOrder, err := o.Service.Create(body)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusCreated, schemas.NewOrder(newOrder))
}

func (o *Orders) Update(ctx *Context) error {
	id, err := ctx.GetUintPathParam("id")
	if err != nil {
		return err
	}
	body := &schemas.UpdateOrder{}
	err = ctx.ParseBody(body)
	if err != nil {
		return err
	}
	updatedOrder, err := o.Service.Update(id, body)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, schemas.NewOrder(updatedOrder))
}

func (o *Orders) Delete(ctx *Context) error {
	id, err := ctx.GetUintPathParam("id")
	if err != nil {
		return err
	}
	err = o.Service.Delete(id)
	if err != nil {
		return err
	}
	return ctx.SendStatus(http.StatusNoContent)
}

func (o *Orders) List(ctx *Context) error {
	orders, err := o.Service.List()
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, schemas.NewOrders(orders))
}

func (o *Orders) Get(ctx *Context) error {
	id, err := ctx.GetUintPathParam("id")
	if err != nil {
		return err
	}
	order, err := o.Service.Get(id)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, schemas.NewOrder(order))
}
