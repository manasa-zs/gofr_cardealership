package db

import (
	"github.com/google/uuid"

	"cardealership/models"
)

var Engine = []models.Engine{
	{ID: uuid.MustParse("045b658e-9160-4f55-8e5a-be8ceb13fbf5"), Displacement: 20, NoOfCylinders: 3, Range: 0},
}
