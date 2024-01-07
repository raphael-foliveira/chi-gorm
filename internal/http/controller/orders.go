package controller

import (
	"net/http"

	"github.com/raphael-foliveira/chi-gorm/internal/http/res"
	"github.com/raphael-foliveira/chi-gorm/internal/http/schemas"
	"github.com/raphael-foliveira/chi-gorm/internal/service"
)

type orders struct {
	service service.Orders
}

func NewOrders(service service.Orders) Controller {
	return &orders{service}
}

func (c *orders) Create(w http.ResponseWriter, r *http.Request) error {
	body, err := parseBody(r, &schemas.CreateOrder{})
	if err != nil {
		return err
	}
	newOrder, err := c.service.Create(body)
	if err != nil {
		return err
	}
	return res.JSON(w, http.StatusCreated, schemas.NewOrder(newOrder))
}

func (c *orders) Update(w http.ResponseWriter, r *http.Request) error {
	id, err := getUintPathParam(r, "id")
	if err != nil {
		return err
	}
	body, err := parseBody(r, &schemas.UpdateOrder{})
	if err != nil {
		return err
	}
	updatedOrder, err := c.service.Update(id, body)
	if err != nil {
		return err
	}
	return res.JSON(w, http.StatusOK, schemas.NewOrder(updatedOrder))
}

func (c *orders) Delete(w http.ResponseWriter, r *http.Request) error {
	id, err := getUintPathParam(r, "id")
	if err != nil {
		return err
	}
	err = c.service.Delete(id)
	if err != nil {
		return err
	}
	return res.SendStatus(w, http.StatusNoContent)
}

func (c *orders) List(w http.ResponseWriter, r *http.Request) error {
	orders, err := c.service.List()
	if err != nil {
		return err
	}
	return res.JSON(w, http.StatusOK, schemas.NewOrders(orders))
}

func (c *orders) Get(w http.ResponseWriter, r *http.Request) error {
	id, err := getUintPathParam(r, "id")
	if err != nil {
		return err
	}
	order, err := c.service.Get(id)
	if err != nil {
		return err
	}
	return res.JSON(w, http.StatusOK, schemas.NewOrder(order))
}
