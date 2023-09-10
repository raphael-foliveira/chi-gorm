package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/raphael-foliveira/chi-gorm/pkg/interfaces"
	"github.com/raphael-foliveira/chi-gorm/pkg/models"
	"github.com/raphael-foliveira/chi-gorm/pkg/res"
	"github.com/raphael-foliveira/chi-gorm/pkg/schemas"
)

type Orders struct {
	repository interfaces.Repository[models.Order]
}

func NewOrders(r interfaces.Repository[models.Order]) *Orders {
	return &Orders{r}
}

func (c *Orders) Create(w http.ResponseWriter, r *http.Request) error {
	var createOrderSchema schemas.CreateOrder
	err := json.NewDecoder(r.Body).Decode(&createOrderSchema)
	if err != nil {
		return res.New(w).Status(http.StatusBadRequest).Error("invalid request body")
	}
	newOrder := models.Order{
		ClientID:  createOrderSchema.ClientID,
		ProductID: createOrderSchema.ProductID,
		Quantity:  createOrderSchema.Quantity,
	}
	err = c.repository.Create(&newOrder)
	if err != nil {
		return res.New(w).Status(http.StatusInternalServerError).Error("internal server error")
	}
	defer r.Body.Close()
	return res.New(w).Status(http.StatusCreated).JSON(&newOrder)
}

func (c *Orders) Update(w http.ResponseWriter, r *http.Request) error {
	id, err := getIdFromPath(r)
	if err != nil {
		return res.New(w).Status(http.StatusBadRequest).Error("invalid id")

	}
	Order, err := c.repository.Get(id)
	if err != nil {
		return res.New(w).Status(http.StatusNotFound).Error("order not found")

	}
	err = json.NewDecoder(r.Body).Decode(&Order)
	if err != nil {
		return res.New(w).Status(http.StatusBadRequest).Error("invalid request body")

	}
	defer r.Body.Close()
	err = c.repository.Update(&Order)
	if err != nil {
		return res.New(w).Status(http.StatusInternalServerError).Error("internal server error")

	}
	return res.New(w).JSON(&Order)
}

func (c *Orders) Delete(w http.ResponseWriter, r *http.Request) error {
	id, err := getIdFromPath(r)
	if err != nil {
		return res.New(w).Status(http.StatusBadRequest).Error("invalid id")

	}
	order, err := c.repository.Get(id)
	if err != nil {
		return res.New(w).Status(http.StatusNotFound).Error("order not found")

	}
	err = c.repository.Delete(&order)
	if err != nil {
		return res.New(w).Status(http.StatusInternalServerError).Error("internal server error")

	}
	w.WriteHeader(http.StatusNoContent)
	return nil
}

func (c *Orders) List(w http.ResponseWriter, r *http.Request) error {
	orders, err := c.repository.List()
	if err != nil {
		return res.New(w).Status(http.StatusInternalServerError).Error("internal server error")
	}
	return res.New(w).JSON(&orders)
}

func (c *Orders) Get(w http.ResponseWriter, r *http.Request) error {
	id, err := getIdFromPath(r)
	if err != nil {
		return res.New(w).Status(http.StatusBadRequest).Error("invalid id")
	}
	order, err := c.repository.Get(id)
	if err != nil {
		return res.New(w).Status(http.StatusNotFound).Error("order not found")
	}
	return res.New(w).JSON(&order)
}
