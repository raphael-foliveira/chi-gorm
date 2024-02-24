package controller

import (
	"net/http"

	"github.com/raphael-foliveira/chi-gorm/internal/http/res"
	"github.com/raphael-foliveira/chi-gorm/internal/http/schemas"
	"github.com/raphael-foliveira/chi-gorm/internal/service"
)

func Orders() *OrdersController {
	return NewOrders(service.Orders())
}

type OrdersController struct {
	service *service.OrdersService
}

func NewOrders(service *service.OrdersService) *OrdersController {
	return &OrdersController{service}
}

func (o *OrdersController) Create(w http.ResponseWriter, r *http.Request) error {
	body, err := parseBody(r, &schemas.CreateOrder{})
	if err != nil {
		return err
	}
	newOrder, err := o.service.Create(body)
	if err != nil {
		return err
	}
	return res.JSON(w, http.StatusCreated, schemas.NewOrder(newOrder))
}

func (o *OrdersController) Update(w http.ResponseWriter, r *http.Request) error {
	id, err := getUintPathParam(r, "id")
	if err != nil {
		return err
	}
	body, err := parseBody(r, &schemas.UpdateOrder{})
	if err != nil {
		return err
	}
	updatedOrder, err := o.service.Update(id, body)
	if err != nil {
		return err
	}
	return res.JSON(w, http.StatusOK, schemas.NewOrder(updatedOrder))
}

func (o *OrdersController) Delete(w http.ResponseWriter, r *http.Request) error {
	id, err := getUintPathParam(r, "id")
	if err != nil {
		return err
	}
	err = o.service.Delete(id)
	if err != nil {
		return err
	}
	return res.SendStatus(w, http.StatusNoContent)
}

func (o *OrdersController) List(w http.ResponseWriter, r *http.Request) error {
	orders, err := o.service.List()
	if err != nil {
		return err
	}
	return res.JSON(w, http.StatusOK, schemas.NewOrders(orders))
}

func (o *OrdersController) Get(w http.ResponseWriter, r *http.Request) error {
	id, err := getUintPathParam(r, "id")
	if err != nil {
		return err
	}
	order, err := o.service.Get(id)
	if err != nil {
		return err
	}
	return res.JSON(w, http.StatusOK, schemas.NewOrder(order))
}
