package models

import "github.com/google/uuid"

type Brand string

const (
	Tesla    Brand = "tesla"
	Porsche  Brand = "porsche"
	Ferrari  Brand = "ferrari"
	Mercedes Brand = "mercedes"
	BMW      Brand = "bmw"
)

type FuelType string

const (
	Electric FuelType = "electric"
	Petrol   FuelType = "petrol"
	Diesel   FuelType = "diesel"
)

type Car struct {
	ID     uuid.UUID `json:"id"`
	Name   string    `json:"name"`
	Year   int       `json:"year"`
	Brand  Brand     `json:"brand"`
	Fuel   FuelType  `json:"fuel"`
	Engine Engine    `json:"engine,omitempty"`
}
