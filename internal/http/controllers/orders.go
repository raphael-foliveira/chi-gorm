package controllers

import (
	"net/http"

	"github.com/raphael-foliveira/chi-gorm/internal/http/res"
	"github.com/raphael-foliveira/chi-gorm/internal/http/schemas"
	"github.com/raphael-foliveira/chi-gorm/internal/repository"
)

type Orders struct {
	repository repository.Orders
}

func NewOrders(ordersRepo repository.Orders) *Orders {
	return &Orders{ordersRepo}
}

func (c *Orders) Create(w http.ResponseWriter, r *http.Request) error {
	var body schemas.CreateOrder
	err := parseBody(r, &body)
	if err != nil {
		return res.Error(w, err, http.StatusBadRequest)
	}
	newOrder := body.ToModel()
	err = c.repository.Create(newOrder)
	if err != nil {
		return res.Error(w, err, http.StatusInternalServerError)
	}
	return res.JSON(w, http.StatusCreated, schemas.NewOrder(newOrder))
}

func (c *Orders) Update(w http.ResponseWriter, r *http.Request) error {
	id, err := getIdFromPath(r)
	if err != nil {
		return res.Error(w, err, http.StatusBadRequest)
	}
	order, err := c.repository.Get(id)
	if err != nil {
		return res.Error(w, err, http.StatusNotFound)
	}
	_, err = c.repository.Get(id)
	if err != nil {
		return res.Error(w, err, http.StatusNotFound)
	}
	var body schemas.UpdateOrder
	err = parseBody(r, &body)
	if err != nil {
		return res.Error(w, err, http.StatusBadRequest)
	}
	order.Quantity = body.Quantity
	err = c.repository.Update(order)
	if err != nil {
		return res.Error(w, err, http.StatusInternalServerError)
	}
	return res.JSON(w, http.StatusOK, schemas.NewOrder(order))
}

func (c *Orders) Delete(w http.ResponseWriter, r *http.Request) error {
	id, err := getIdFromPath(r)
	if err != nil {
		return res.Error(w, err, http.StatusBadRequest)
	}
	order, err := c.repository.Get(id)
	if err != nil {
		return res.Error(w, err, http.StatusNotFound)
	}
	err = c.repository.Delete(order)
	if err != nil {
		return res.Error(w, err, http.StatusInternalServerError)
	}
	return res.SendStatus(w, http.StatusNoContent)
}

func (c *Orders) List(w http.ResponseWriter, r *http.Request) error {
	orders, err := c.repository.List()
	if err != nil {
		return res.Error(w, err, http.StatusInternalServerError)
	}
	return res.JSON(w, http.StatusOK, schemas.NewOrders(orders))
}

func (c *Orders) Get(w http.ResponseWriter, r *http.Request) error {
	id, err := getIdFromPath(r)
	if err != nil {
		return res.Error(w, err, http.StatusBadRequest)
	}
	order, err := c.repository.Get(id)
	if err != nil {
		return res.Error(w, err, http.StatusNotFound)
	}
	return res.JSON(w, http.StatusOK, schemas.NewOrder(order))
}
