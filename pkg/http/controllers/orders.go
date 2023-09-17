package controllers

import (
	"net/http"

	"github.com/raphael-foliveira/chi-gorm/pkg/http/res"
	"github.com/raphael-foliveira/chi-gorm/pkg/http/schemas"
	"github.com/raphael-foliveira/chi-gorm/pkg/models"
	"github.com/raphael-foliveira/chi-gorm/pkg/persistence/store"
)

type Orders struct {
	repository store.Orders
}

func NewOrders(r store.Orders) *Orders {
	return &Orders{r}
}

func (c *Orders) Create(w http.ResponseWriter, r *http.Request) error {
	var body schemas.CreateOrder
	err := parseBody(r, &body)
	if err != nil {
		return res.Error(w, err, http.StatusBadRequest, "bad request")
	}
	newOrder := models.Order{
		ClientID:  body.ClientID,
		ProductID: body.ProductID,
		Quantity:  body.Quantity,
	}
	err = c.repository.Create(&newOrder)
	if err != nil {
		return res.Error(w, err, http.StatusInternalServerError, "internal server error")
	}
	return res.JSON(w, http.StatusCreated, &newOrder)
}

func (c *Orders) Update(w http.ResponseWriter, r *http.Request) error {
	id, err := getIdFromPath(r)
	if err != nil {
		return res.Error(w, err, http.StatusBadRequest, err.Error())
	}
	order, err := c.repository.Get(id)
	if err != nil {
		return res.Error(w, err, http.StatusNotFound, "order not found")
	}
	var body schemas.UpdateOrder
	err = parseBody(r, &body)
	if err != nil {
		return res.Error(w, err, http.StatusBadRequest, "bad request")
	}
	order.Quantity = body.Quantity
	err = c.repository.Update(order)
	if err != nil {
		return res.Error(w, err, http.StatusInternalServerError, "internal server error")
	}
	return res.JSON(w, http.StatusOK, &order)
}

func (c *Orders) Delete(w http.ResponseWriter, r *http.Request) error {
	id, err := getIdFromPath(r)
	if err != nil {
		return res.Error(w, err, http.StatusBadRequest, err.Error())

	}
	order, err := c.repository.Get(id)
	if err != nil {
		return res.Error(w, err, http.StatusNotFound, "order not found")

	}
	err = c.repository.Delete(order)
	if err != nil {
		return res.Error(w, err, http.StatusInternalServerError, "internal server error")

	}
	return res.SendStatus(w, http.StatusNoContent)
}

func (c *Orders) List(w http.ResponseWriter, r *http.Request) error {
	orders, err := c.repository.List()
	if err != nil {
		return res.Error(w, err, http.StatusInternalServerError, "internal server error")
	}
	return res.JSON(w, http.StatusOK, &orders)
}

func (c *Orders) Get(w http.ResponseWriter, r *http.Request) error {
	id, err := getIdFromPath(r)
	if err != nil {
		return res.Error(w, err, http.StatusBadRequest, err.Error())
	}
	order, err := c.repository.Get(id)
	if err != nil {
		return res.Error(w, err, http.StatusNotFound, "order not found")
	}
	return res.JSON(w, http.StatusOK, &order)
}
