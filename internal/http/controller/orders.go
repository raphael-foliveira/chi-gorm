package controller

import (
	"net/http"

	"github.com/raphael-foliveira/chi-gorm/internal/http/schemas"
	"github.com/raphael-foliveira/chi-gorm/internal/service"
)

type orders struct{}

func NewOrders() *orders {
	return &orders{}
}

func (o *orders) Create(ctx *Context) error {
	body := &schemas.CreateOrder{}
	err := ctx.ParseBody(body)
	if err != nil {
		return err
	}
	newOrder, err := service.Orders.Create(body)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusCreated, schemas.NewOrder(newOrder))
}

func (o *orders) Update(ctx *Context) error {
	id, err := ctx.GetUintPathParam("id")
	if err != nil {
		return err
	}
	body := &schemas.UpdateOrder{}
	err = ctx.ParseBody(body)
	if err != nil {
		return err
	}
	updatedOrder, err := service.Orders.Update(id, body)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, schemas.NewOrder(updatedOrder))
}

func (o *orders) Delete(ctx *Context) error {
	id, err := ctx.GetUintPathParam("id")
	if err != nil {
		return err
	}
	err = service.Orders.Delete(id)
	if err != nil {
		return err
	}
	return ctx.SendStatus(http.StatusNoContent)
}

func (o *orders) List(ctx *Context) error {
	orders, err := service.Orders.List()
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, schemas.NewOrders(orders))
}

func (o *orders) Get(ctx *Context) error {
	id, err := ctx.GetUintPathParam("id")
	if err != nil {
		return err
	}
	order, err := service.Orders.Get(id)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, schemas.NewOrder(order))
}
