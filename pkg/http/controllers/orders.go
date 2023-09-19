package controllers

import (
	"net/http"

	"github.com/raphael-foliveira/chi-gorm/pkg/http/res"
	"github.com/raphael-foliveira/chi-gorm/pkg/http/schemas"
	"github.com/raphael-foliveira/chi-gorm/pkg/repository"
)

type Orders struct {
	ordersRepo   repository.Orders
	clientsRepo  repository.Clients
	productsRepo repository.Products
}

func NewOrders(ordersStore repository.Orders, clientsStore repository.Clients, productsStore repository.Products) *Orders {
	return &Orders{ordersStore, clientsStore, productsStore}
}

func (c *Orders) Create(w http.ResponseWriter, r *http.Request) error {
	var body schemas.CreateOrder
	err := parseBody(r, &body)
	if err != nil {
		return res.Error(w, err, http.StatusBadRequest, "bad request")
	}
	newOrder := body.ToModel()
	err = c.ordersRepo.Create(newOrder)
	if err != nil {
		return res.Error(w, err, http.StatusInternalServerError, "internal server error")
	}
	return res.JSON(w, http.StatusCreated, schemas.NewOrder(newOrder))
}

func (c *Orders) Update(w http.ResponseWriter, r *http.Request) error {
	id, err := getIdFromPath(r)
	if err != nil {
		return res.Error(w, err, http.StatusBadRequest, err.Error())
	}
	order, err := c.ordersRepo.Get(id)
	if err != nil {
		return res.Error(w, err, http.StatusNotFound, "order not found")
	}
	var body schemas.UpdateOrder
	err = parseBody(r, &body)
	if err != nil {
		return res.Error(w, err, http.StatusBadRequest, "bad request")
	}
	order.Quantity = body.Quantity
	err = c.ordersRepo.Update(order)
	if err != nil {
		return res.Error(w, err, http.StatusInternalServerError, "internal server error")
	}
	return res.JSON(w, http.StatusOK, schemas.NewOrder(order))
}

func (c *Orders) Delete(w http.ResponseWriter, r *http.Request) error {
	id, err := getIdFromPath(r)
	if err != nil {
		return res.Error(w, err, http.StatusBadRequest, err.Error())
	}
	order, err := c.ordersRepo.Get(id)
	if err != nil {
		return res.Error(w, err, http.StatusNotFound, "order not found")
	}
	err = c.ordersRepo.Delete(order)
	if err != nil {
		return res.Error(w, err, http.StatusInternalServerError, "internal server error")
	}
	return res.SendStatus(w, http.StatusNoContent)
}

func (c *Orders) List(w http.ResponseWriter, r *http.Request) error {
	orders, err := c.ordersRepo.List()
	if err != nil {
		return res.Error(w, err, http.StatusInternalServerError, "internal server error")
	}
	return res.JSON(w, http.StatusOK, schemas.NewOrders(orders))
}

func (c *Orders) Get(w http.ResponseWriter, r *http.Request) error {
	id, err := getIdFromPath(r)
	if err != nil {
		return res.Error(w, err, http.StatusBadRequest, err.Error())
	}
	order, err := c.ordersRepo.Get(id)
	if err != nil {
		return res.Error(w, err, http.StatusNotFound, "order not found")
	}
	client, err := c.clientsRepo.Get(order.ClientID)
	if err != nil {
		return res.Error(w, err, http.StatusNotFound, "client not found")
	}
	product, err := c.productsRepo.Get(order.ProductID)
	if err != nil {
		return res.Error(w, err, http.StatusNotFound, "product not found")
	}
	return res.JSON(w, http.StatusOK, schemas.NewOrderDetail(order, client, product))
}
