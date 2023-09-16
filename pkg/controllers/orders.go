package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/raphael-foliveira/chi-gorm/pkg/models"
	"github.com/raphael-foliveira/chi-gorm/pkg/repositories"
	"github.com/raphael-foliveira/chi-gorm/pkg/res"
	"github.com/raphael-foliveira/chi-gorm/pkg/schemas"
)

type Orders struct {
	repository repositories.Orders
}

func NewOrders(r repositories.Orders) *Orders {
	return &Orders{r}
}

func (c *Orders) Create(w http.ResponseWriter, r *http.Request) error {
	body, err := c.parseCreate(w, r)
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
		return res.Error(w, err, http.StatusBadRequest, "invalid id")
	}
	order, err := c.repository.Get(id)
	if err != nil {
		return res.Error(w, err, http.StatusNotFound, "order not found")
	}
	body, err := c.parseUpdate(w, r)
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
		return res.Error(w, err, http.StatusBadRequest, "invalid id")

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
		return res.Error(w, err, http.StatusBadRequest, "invalid id")
	}
	order, err := c.repository.Get(id)
	if err != nil {
		return res.Error(w, err, http.StatusNotFound, "order not found")
	}
	return res.JSON(w, http.StatusOK, &order)
}

func (c *Orders) parseCreate(w http.ResponseWriter, r *http.Request) (*schemas.CreateOrder, error) {
	defer r.Body.Close()
	body := schemas.CreateOrder{}
	return &body, json.NewDecoder(r.Body).Decode(&body)
}

func (c *Orders) parseUpdate(w http.ResponseWriter, r *http.Request) (*schemas.UpdateOrder, error) {
	defer r.Body.Close()
	body := schemas.UpdateOrder{}
	return &body, json.NewDecoder(r.Body).Decode(&body)
}
