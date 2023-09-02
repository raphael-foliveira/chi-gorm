package product

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/raphael-foliveira/chi-gorm/internal/interfaces"
	"github.com/raphael-foliveira/chi-gorm/pkg/res"
)

type Controller struct {
	repository interfaces.IRepository[Product]
}

func NewController(r interfaces.IRepository[Product]) *Controller {
	return &Controller{r}
}

func (c *Controller) Create(w http.ResponseWriter, r *http.Request) error {
	var newProduct Product
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

func (c *Controller) Update(w http.ResponseWriter, r *http.Request) error {
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

func (c *Controller) Delete(w http.ResponseWriter, r *http.Request) error {
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

func (c *Controller) List(w http.ResponseWriter, r *http.Request) error {
	products, err := c.repository.List()
	if err != nil {
		return res.Error(w, http.StatusInternalServerError, "internal server error", err)
	}
	return res.JSON(w, http.StatusOK, &products)
}

func (c *Controller) Get(w http.ResponseWriter, r *http.Request) error {
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
