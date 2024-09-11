package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/raphael-foliveira/chi-gorm/internal/exceptions"
	"github.com/raphael-foliveira/chi-gorm/internal/http/schemas"
	"github.com/raphael-foliveira/chi-gorm/internal/validate"
)

var (
	Clients     = NewClients()
	HealthCheck = NewHealthCheck()
	Orders      = NewOrders()
	Products    = NewProducts()
)

type Context struct {
	Response http.ResponseWriter
	Request  *http.Request
}

func NewContext(w http.ResponseWriter, r *http.Request) *Context {
	return &Context{
		Response: w,
		Request:  r,
	}
}

func (c *Context) SendStatus(status int) error {
	c.Response.WriteHeader(status)
	return nil
}

func (c *Context) JSON(status int, data interface{}) error {
	c.Response.Header().Set("Content-Type", "application/json")
	c.Response.WriteHeader(status)
	return json.NewEncoder(c.Response).Encode(data)
}

func (c *Context) GetUintPathParam(paramName string) (uint, error) {
	id, err := strconv.ParseUint(chi.URLParam(c.Request, paramName), 10, 64)
	if err != nil {
		return 0, exceptions.BadRequest(fmt.Sprintf("invalid %v", paramName))
	}
	return uint(id), nil
}

func (c *Context) ParseBody(v schemas.Validatable) error {
	err := json.NewDecoder(c.Request.Body).Decode(v)
	defer c.Request.Body.Close()
	if err != nil {
		return exceptions.UnprocessableEntity("invalid body")
	}
	if err := v.Validate(); err != nil {
		return exceptions.NewApiError(http.StatusUnprocessableEntity, err)
	}
	return nil
}

func useHandler(fn ControllerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		context := NewContext(w, r)
		err := fn(context)
		if err != nil {
			handleApiErr(context, err)
		}
	}
}

func handleApiErr(ctx *Context, err error) error {
	slog.Error(err.Error())
	apiErr := &exceptions.ApiError{
		Status: http.StatusInternalServerError,
		Err:    "internal server error",
	}
	validationErr := &validate.ValidationError{}
	if errors.As(err, &validationErr) {
		return ctx.JSON(http.StatusUnprocessableEntity, validationErr)
	}
	errors.As(err, &apiErr)
	return ctx.JSON(apiErr.Status, apiErr)
}

type ControllerFunc func(*Context) error
