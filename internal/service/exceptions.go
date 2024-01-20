package service

import "github.com/raphael-foliveira/chi-gorm/internal/exceptions"

var errClientNotFound = exceptions.NotFound("client not found")
var errOrderNotFound = exceptions.NotFound("order not found")
var errProductNotFound = exceptions.NotFound("product not found")
