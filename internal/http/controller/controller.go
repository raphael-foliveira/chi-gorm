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
	"github.com/raphael-foliveira/chi-gorm/internal/validate"
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

func (c *Context) JSON(status int, data any) error {
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

type Validatable interface {
	Validate() error
}

func (c *Context) ParseBody(v Validatable) error {
	err := json.NewDecoder(c.Request.Body).Decode(v)
	if err != nil {
		return exceptions.UnprocessableEntity("invalid body")
	}
	defer c.Request.Body.Close()
	if err := v.Validate(); err != nil {
		return exceptions.NewApiError(http.StatusUnprocessableEntity, err)
	}
	return nil
}

func useHandler(fn ControllerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		context := NewContext(w, r)
		if err := fn(context); err != nil {
			if err := handleApiErr(context, err); err != nil {
				slog.Error(
					"error handling api error",
					slog.Any("error", err),
				)
			}
		}
	}
}

func handleApiErr(ctx *Context, err error) error {
	apiErr := &exceptions.ApiError{
		Status: http.StatusInternalServerError,
		Err:    "internal server error",
	}

	validationErr := &validate.ValidationError{}
	if errors.As(err, &validationErr) {
		return ctx.JSON(http.StatusUnprocessableEntity, validationErr)
	}

	// log unhandled errors
	if !errors.As(err, &apiErr) {
		slog.Error(err.Error())
	}

	return ctx.JSON(apiErr.Status, apiErr)
}

type ControllerFunc func(*Context) error

type Router struct {
	*chi.Mux
}

func NewRouter() *Router {
	return &Router{Mux: chi.NewRouter()}
}

func (r *Router) Get(path string, controller ControllerFunc) {
	r.Mux.Get(path, useHandler(controller))
}

func (r *Router) Post(path string, controller ControllerFunc) {
	r.Mux.Post(path, useHandler(controller))
}

func (r *Router) Patch(path string, controller ControllerFunc) {
	r.Mux.Patch(path, useHandler(controller))
}

func (r *Router) Delete(path string, controller ControllerFunc) {
	r.Mux.Delete(path, useHandler(controller))
}

func (r *Router) Put(path string, controller ControllerFunc) {
	r.Mux.Put(path, useHandler(controller))
}
