package service

import "github.com/raphael-foliveira/chi-gorm/internal/exceptions"

var (
	errClientNotFound  = exceptions.NotFound("client not found")
	errOrderNotFound   = exceptions.NotFound("order not found")
	errProductNotFound = exceptions.NotFound("product not found")
)
