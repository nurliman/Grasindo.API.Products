package main

import (
	"context"

	"github.com/nurliman/Grasindo.API.Products/models"
)

const (
	apiVersion = "v1"
)

// Service interface
type Service interface {
	CreateProduct(ctx context.Context, product models.Product) (int, error)
}

// newService Create new service instance
func newService() Service {
	return service{}
}

// service implementation
type service struct{}

func (service) CreateProduct(ctx context.Context, product models.Product) (int, error) {
	return 0, nil
}
