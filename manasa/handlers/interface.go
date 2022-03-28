package handlers

import (
	"context"

	"github.com/google/uuid"

	"manasa/models"
)

type Service interface {
	GetByID(ctx context.Context, id uuid.UUID) (*models.Car, error)
	GetByBrand(ctx context.Context, brand string, isEngine bool) ([]*models.Car, error)
	Create(ctx context.Context, car *models.Car) (*models.Car, error)
	Delete(ctx context.Context, id uuid.UUID) error
	Update(ctx context.Context, car *models.Car) (*models.Car, error)
}
