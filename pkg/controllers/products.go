package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/raphael-foliveira/chi-gorm/pkg/interfaces"
	"github.com/raphael-foliveira/chi-gorm/pkg/models"
	"github.com/raphael-foliveira/chi-gorm/pkg/res"
)

type ProductsController struct {
	repository interfaces.Repository[models.Product]
}

func NewProductsController(r interfaces.Repository[models.Product]) *ProductsController {
	return &ProductsController{r}
}

func (c *ProductsController) Create(w http.ResponseWriter, r *http.Request) error {
	var newProduct models.Product
	err := json.NewDecoder(r.Body).Decode(&newProduct)
	if err != nil {
		return res.Error(w, http.StatusBadRequest, "bad request", err)
	}
	defer r.Body.Close()
	err = c.repository.Create(&newProduct)
	if err != nil {
		return res.Error(w, http.StatusInternalServerError, "internal server error", err)
	}
	return res.JSON(w, http.StatusCreated, &newProduct)
}

func (c *ProductsController) Update(w http.ResponseWriter, r *http.Request) error {
	id, err := strconv.ParseUint(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		return res.Error(w, http.StatusBadRequest, "bad request", err)
	}
	product, err := c.repository.Get(id)
	if err != nil {
		return res.Error(w, http.StatusNotFound, "product not found", err)
	}
	err = json.NewDecoder(r.Body).Decode(&product)
	if err != nil {
		return res.Error(w, http.StatusBadRequest, "bad request", err)
	}
	defer r.Body.Close()
	err = c.repository.Update(&product)
	if err != nil {
		return res.Error(w, http.StatusInternalServerError, "internal server error", err)
	}
	return res.JSON(w, http.StatusOK, &product)
}

func (c *ProductsController) Delete(w http.ResponseWriter, r *http.Request) error {
	id, err := strconv.ParseUint(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		return res.Error(w, http.StatusBadRequest, "bad request", err)
	}
	product, err := c.repository.Get(id)
	if err != nil {
		return res.Error(w, http.StatusNotFound, "product not found", err)
	}
	err = c.repository.Delete(&product)
	if err != nil {
		return res.Error(w, http.StatusInternalServerError, "internal server error", err)
	}
	w.WriteHeader(http.StatusNoContent)
	return nil
}

func (c *ProductsController) List(w http.ResponseWriter, r *http.Request) error {
	products, err := c.repository.List()
	if err != nil {
		return res.Error(w, http.StatusInternalServerError, "internal server error", err)
	}
	return res.JSON(w, http.StatusOK, &products)
}

func (c *ProductsController) Get(w http.ResponseWriter, r *http.Request) error {
	id, err := strconv.ParseUint(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		return res.Error(w, http.StatusBadRequest, "bad request", err)
	}
	product, err := c.repository.Get(id)
	if err != nil {
		return res.Error(w, http.StatusNotFound, "product not found", err)
	}
	return res.JSON(w, http.StatusOK, &product)
}
