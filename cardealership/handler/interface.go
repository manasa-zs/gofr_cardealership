package handler

import (
	"cardealership/models"
	"developer.zopsmart.com/go/gofr/pkg/gofr"
	"github.com/google/uuid"
)

type Service interface {
	Create(*gofr.Context, *models.Car) (*models.Car, error)
	GetByID(*gofr.Context, uuid.UUID) (*models.Car, error)
	GetByBrand(*gofr.Context, string, bool) ([]*models.Car, error)
	Delete(*gofr.Context, uuid.UUID) error
	Update(*gofr.Context, *models.Car) (*models.Car, error)
}
