package services

import (
	"context"

	"github.com/google/uuid"

	"manasa/models"
)

type Car interface {
	CarCreate(context.Context, *models.Car) (*models.Car, error)
	CarGetByID(context.Context, uuid.UUID) (*models.Car, error)
	CarGetByBrand(context.Context, string) ([]*models.Car, error)
	CarUpdate(context.Context, *models.Car) (*models.Car, error)
	CarDelete(context.Context, uuid.UUID) error
}

type Engine interface {
	EngineCreate(context.Context, *models.Engine) (*models.Engine, error)
	EngineGetByID(context.Context, uuid.UUID) (*models.Engine, error)
	EngineUpdate(context.Context, *models.Engine) (*models.Engine, error)
	EngineDelete(context.Context, uuid.UUID) error
}
