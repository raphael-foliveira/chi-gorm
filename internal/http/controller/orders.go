package controller

import (
	"net/http"

	"github.com/raphael-foliveira/chi-gorm/internal/http/res"
	"github.com/raphael-foliveira/chi-gorm/internal/http/schemas"
	"github.com/raphael-foliveira/chi-gorm/internal/service"
)

var Orders = &orders{}

type orders struct{}

func (c *orders) Create(w http.ResponseWriter, r *http.Request) error {
	body, err := parseBody(r, &schemas.CreateOrder{})
	if err != nil {
		return err
	}
	newOrder, err := service.Orders.Create(body)
	if err != nil {
		return err
	}
	return res.JSON(w, http.StatusCreated, schemas.NewOrder(newOrder))
}

func (c *orders) Update(w http.ResponseWriter, r *http.Request) error {
	id, err := getIdFromPath(r)
	if err != nil {
		return err
	}
	body, err := parseBody(r, &schemas.UpdateOrder{})
	if err != nil {
		return err
	}
	updatedOrder, err := service.Orders.Update(id, body)
	if err != nil {
		return err
	}
	return res.JSON(w, http.StatusOK, schemas.NewOrder(updatedOrder))
}

func (c *orders) Delete(w http.ResponseWriter, r *http.Request) error {
	id, err := getIdFromPath(r)
	if err != nil {
		return err
	}
	err = service.Orders.Delete(id)
	if err != nil {
		return err
	}
	return res.SendStatus(w, http.StatusNoContent)
}

func (c *orders) List(w http.ResponseWriter, r *http.Request) error {
	orders, err := service.Orders.List()
	if err != nil {
		return err
	}
	return res.JSON(w, http.StatusOK, schemas.NewOrders(orders))
}

func (c *orders) Get(w http.ResponseWriter, r *http.Request) error {
	id, err := getIdFromPath(r)
	if err != nil {
		return err
	}
	order, err := service.Orders.Get(id)
	if err != nil {
		return err
	}
	return res.JSON(w, http.StatusOK, schemas.NewOrder(order))
}
