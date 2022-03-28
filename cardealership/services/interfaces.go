package services

import (
	"cardealership/models"
	"developer.zopsmart.com/go/gofr/pkg/gofr"
	"github.com/google/uuid"
)

type Car interface {
	CarCreate(*gofr.Context, *models.Car) (*models.Car, error)
	CarGet(*gofr.Context, uuid.UUID) (*models.Car, error)
	CarGetByBrand(*gofr.Context, string, bool) ([]*models.Car, error)
	CarDelete(*gofr.Context, uuid.UUID) error
	CarUpdate(*gofr.Context, *models.Car) (*models.Car, error)
}

type Engine interface {
	EngineCreate(*gofr.Context, *models.Engine) (*models.Engine, error)
	EngineGet(*gofr.Context, uuid.UUID) (*models.Engine, error)
	EngineDelete(*gofr.Context, uuid.UUID) error
	EngineUpdate(*gofr.Context, *models.Engine) (*models.Engine, error)
}
