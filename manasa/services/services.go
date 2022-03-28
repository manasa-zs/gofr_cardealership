package services

import (
	"context"
	"errors"

	"github.com/google/uuid"

	"manasa/models"
)

type Service struct {
	car    Car
	engine Engine
}

// New is a factory function in Service layer
func New(c Car, e Engine) Service {
	return Service{car: c, engine: e}
}

// GetByID method is used in service layer to fetch rows.
func (s Service) GetByID(ctx context.Context, id uuid.UUID) (*models.Car, error) {
	res, err := s.car.CarGetByID(ctx, id)
	if err != nil {
		return &models.Car{}, err
	}

	res.ID = id

	engine, err := s.engine.EngineGetByID(ctx, res.Engine.ID)
	if err != nil {
		return &models.Car{}, err
	}

	res.Engine = *engine

	return res, nil
}

// GetByBrand method is used in service layer to fetch rows.
func (s Service) GetByBrand(ctx context.Context, brand string, isEngine bool) ([]*models.Car, error) { //nolint
	if brand != "bmw" && brand != "tesla" && brand != "ferrari" && brand != "mercedes" && brand != "porsche" {
		return []*models.Car{}, errors.New("invalid brand")
	}

	res, err := s.car.CarGetByBrand(ctx, brand)
	if err != nil {
		return []*models.Car{}, err
	}

	if isEngine {
		for i := 0; i < len(res); i++ {
			engine, er := s.engine.EngineGetByID(ctx, res[i].Engine.ID)
			if er != nil {
				return []*models.Car{}, er
			}

			res[i].Engine = *engine
		}
	}

	return res, err
}

// Create method is used in service layer for inserting rows.
func (s Service) Create(ctx context.Context, car *models.Car) (*models.Car, error) {
	err := validation(*car)
	if err != nil {
		return &models.Car{}, err
	}

	resp, err := s.car.CarCreate(ctx, car)
	if err != nil {
		return &models.Car{}, err
	}

	car.Engine.ID = resp.Engine.ID

	engine, err := s.engine.EngineCreate(ctx, &car.Engine)
	if err != nil {
		return &models.Car{}, err
	}

	resp.Engine = *engine

	return resp, nil
}

// Delete method is used in service layer to delete rows.
func (s Service) Delete(ctx context.Context, id uuid.UUID) error {
	car, err := s.car.CarGetByID(ctx, id)
	if err != nil {
		return err
	}

	err = s.engine.EngineDelete(ctx, car.Engine.ID)
	if err != nil {
		return err
	}

	err = s.car.CarDelete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

// Update method is used in Service layer to update rows.
func (s Service) Update(ctx context.Context, car *models.Car) (*models.Car, error) {
	err := validation(*car)
	if err != nil {
		return &models.Car{}, err
	}

	resp, err := s.car.CarUpdate(ctx, car)

	if err != nil {
		return &models.Car{}, err
	}

	res, err := s.car.CarGetByID(ctx, resp.ID)
	if err != nil {
		return &models.Car{}, err
	}

	car.Engine.ID = res.Engine.ID

	engine, err := s.engine.EngineUpdate(ctx, &car.Engine)
	if err != nil {
		return &models.Car{}, err
	}

	resp.Engine = *engine

	return resp, nil
}

// validation function is used in service layer for validation of inputs provided.
func validation(car models.Car) error { //nolint
	if car.Year > 2022 || car.Year < 1980 {
		return errors.New("invalid Year")
	}

	if car.Brand != models.BMW && car.Brand != models.Tesla && car.Brand != models.Ferrari &&
		car.Brand != models.Mercedes && car.Brand != models.Porsche {
		return errors.New("invalid brand")
	} else if car.Fuel != models.Petrol && car.Fuel != models.Diesel && car.Fuel != models.Electric {
		return errors.New("invalid fuel")
	} else if car.Fuel == models.Petrol && car.Engine.Range != 0 {
		return errors.New("petrol car doesn't have range")
	} else if car.Fuel == models.Diesel && car.Engine.Range != 0 {
		return errors.New("diesel car doesn't have range")
	} else if car.Fuel == models.Electric && (car.Engine.Displacement != 0 || car.Engine.NoOfCylinders != 0) {
		return errors.New("electric car doesn't have displacement and no.of.cylinders")
	}

	return nil
}
