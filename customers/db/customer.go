package db

import (
	"github.com/google/uuid"
	"training/customers/models"
)

var Customer = []models.Customer{
	{ID: uuid.New(), Name: "manasa", Age: 21},
}
