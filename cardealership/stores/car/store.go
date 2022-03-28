package car

import (
	"cardealership/db"
	"cardealership/models"
	"database/sql"
	_ "database/sql"
	"developer.zopsmart.com/go/gofr/pkg/errors"
	"developer.zopsmart.com/go/gofr/pkg/gofr"

	"github.com/google/uuid"
)

type car struct {
	db []models.Car
}

func New() *car {
	return &car{db.Car}
}

// CarCreate method is used to insert a row into the car table
func (c *car) CarCreate(ctx *gofr.Context, car *models.Car) (*models.Car, error) {
	car.ID = uuid.New()
	car.Engine.ID = uuid.New()
	_, err := ctx.DB().ExecContext(ctx, "INSERT INTO car(id,name,year,brand,fueltype,engine_id)"+
		"VALUES(?,?,?,?,?,?)", car.ID, car.Name, car.Year, car.Brand, car.Fuel, car.Engine.ID)

	if err != nil {
		return &models.Car{}, errors.DB{Err: err}
	}
	return car, nil
}

// CarGet is used to get a row of a car based on given id
func (c car) CarGet(ctx *gofr.Context, id uuid.UUID) (*models.Car, error) {
	var car = &models.Car{}

	rows := ctx.DB().QueryRowContext(ctx, "SELECT * FROM car where id=?", id)
	err := rows.Scan(&car.ID, &car.Name, &car.Year, &car.Brand, &car.Fuel, &car.Engine.ID)

	if err != nil {
		return &models.Car{}, errors.DB{Err: err}
	}

	return car, nil
}

// CarUpdate method is used to update/modify a particular row in car table
func (c car) CarUpdate(ctx *gofr.Context, car *models.Car) (*models.Car, error) {
	_, err := ctx.DB().ExecContext(ctx, "UPDATE car SET name=?,year=?,brand=?,fueltype=? "+
		"where id=?", car.Name, car.Year, car.Brand, car.Fuel, car.ID)

	if err != nil {
		return &models.Car{}, errors.DB{Err: err}
	}

	return car, nil
}

// CarDelete method is used to delete a row in car table based on given id
func (c car) CarDelete(ctx *gofr.Context, id uuid.UUID) error {
	_, err := ctx.DB().ExecContext(ctx, "delete from car where id=?", id)

	if err != nil {
		return errors.DB{Err: err}
	}

	return nil
}

// CarGetByBrand method takes brand as input and returns rows with the given brand
func (c car) CarGetByBrand(ctx *gofr.Context, brand string, isEngine bool) ([]*models.Car, error) {
	rows, err := ctx.DB().QueryContext(ctx, "select * from car where brand=?", brand)
	if err != nil {
		return nil, errors.DB{Err: err}
	}

	if rows.Err() != nil {
		return []*models.Car{}, errors.DB{Err: err}
	}

	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)

	var res []*models.Car

	for rows.Next() {
		var car models.Car

		err := rows.Scan(&car.ID, &car.Name, &car.Year, &car.Brand, &car.Fuel, &car.Engine.ID)
		if err != nil {
			return []*models.Car{}, errors.DB{Err: err}
		}

		res = append(res, &car)
	}

	return res, nil
}
