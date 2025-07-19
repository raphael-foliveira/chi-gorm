package api

import (
	"net/http"

	"github.com/raphael-foliveira/chi-gorm/internal/http/schemas"
	"github.com/raphael-foliveira/chi-gorm/internal/ports"
)

type OrdersController struct {
	ordersRepo ports.OrdersRepository
}

func NewOrdersController(ordersRepo ports.OrdersRepository) *OrdersController {
	return &OrdersController{
		ordersRepo: ordersRepo,
	}
}

func (o *OrdersController) Create(ctx *Context) error {
	var body schemas.CreateOrder
	err := ctx.ParseBody(&body)
	if err != nil {
		return err
	}
	newOrder := body.ToModel()
	if err := o.ordersRepo.Create(newOrder); err != nil {
		return err
	}
	return ctx.JSON(http.StatusCreated, schemas.NewOrder(newOrder))
}

func (o *OrdersController) Update(ctx *Context) error {
	id, err := ctx.GetUintPathParam("id")
	if err != nil {
		return err
	}
	var body schemas.UpdateOrder
	err = ctx.ParseBody(&body)
	if err != nil {
		return err
	}
	order := body.ToModel()
	order.ID = id
	if err := o.ordersRepo.Update(order); err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, schemas.NewOrder(order))
}

func (o *OrdersController) Delete(ctx *Context) error {
	id, err := ctx.GetUintPathParam("id")
	if err != nil {
		return err
	}
	err = o.ordersRepo.Delete(id)
	if err != nil {
		return err
	}
	return ctx.SendStatus(http.StatusNoContent)
}

func (o *OrdersController) List(ctx *Context) error {
	orders, err := o.ordersRepo.List()
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, schemas.NewOrders(orders))
}

func (o *OrdersController) Get(ctx *Context) error {
	id, err := ctx.GetUintPathParam("id")
	if err != nil {
		return err
	}
	order, err := o.ordersRepo.Get(id)
	if err != nil {
		return err
	}
	return ctx.JSON(http.StatusOK, schemas.NewOrder(order))
}
