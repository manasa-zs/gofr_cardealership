package handlers

import (
	"developer.zopsmart.com/go/gofr/pkg/gofr"
	"github.com/google/uuid"
	"training/customers/models"
)

type CustomerService interface {
	Create(ctx *gofr.Context, customer *models.Customer) (*models.Customer, error)
	Get(ctx *gofr.Context, id uuid.UUID) (*models.Customer, error)
	Update(ctx *gofr.Context, customer *models.Customer) (*models.Customer, error)
	Delete(ctx *gofr.Context, ID uuid.UUID) error
}
