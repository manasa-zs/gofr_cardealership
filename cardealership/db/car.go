package db

import (
	"github.com/google/uuid"

	"cardealership/models"
)

var Car = []models.Car{
	{ID: uuid.New(), Name: "tesla 1", Year: 2020, Brand: "tesla", Fuel: "diesel"},
}
