package models

import "github.com/google/uuid"

type Engine struct {
	ID            uuid.UUID `json:"id"`
	Displacement  float64   `json:"displacement,omitempty"`
	NoOfCylinders int       `json:"noOfCylinders,omitempty"`
	Range         float64   `json:"range,omitempty"`
}
