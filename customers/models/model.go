package models

import "github.com/google/uuid"

type Customer struct {
	ID   uuid.UUID
	Name string
	Age  int
}
