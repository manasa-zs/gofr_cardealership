package car

import (
	"context"
	"database/sql"
	"errors"

	"github.com/google/uuid"

	"manasa/models"
)

type DB struct {
	DB *sql.DB
}

func New(db *sql.DB) DB {
	return DB{DB: db}
}

const (
	create     = "insert into car (id,name,year,brand,fueltype,engine_id) VALUES (?,?,?,?,?,?)"
	getByID    = "select id,name,year,brand,fueltype,engine_id from car where id = ?"
	getByBrand = "select * from car where brand = ?"
	update     = "update car set name=?,year=?,brand=?,fueltype=? where id=?"
	del        = "delete from car where id = ?"
)

// CarGetByID method used to get values from car table.
func (d DB) CarGetByID(ctx context.Context, id uuid.UUID) (*models.Car, error) {
	var car models.Car

	row := d.DB.QueryRowContext(ctx, getByID, id)

	err := row.Scan(&car.ID, &car.Name, &car.Year, &car.Brand, &car.Fuel, &car.Engine.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			return &models.Car{}, errors.New("entity not found")
		}

		return &models.Car{}, err
	}

	return &car, nil
}

// CarGetByBrand method used to get values from car table based on the brand.
func (d DB) CarGetByBrand(ctx context.Context, brand string) ([]*models.Car, error) {
	rows, err := d.DB.QueryContext(ctx, getByBrand, brand)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("entity not found")
		}

		return nil, errors.New("query error")
	}

	if rows.Err() != nil {
		return nil, rows.Err()
	}

	var cars []*models.Car

	for rows.Next() {
		var car models.Car

		err = rows.Scan(&car.ID, &car.Name, &car.Year, &car.Brand, &car.Fuel, &car.Engine.ID)
		if err != nil {
			return nil, errors.New("error while scanning")
		}

		cars = append(cars, &car)
	}

	defer rows.Close()

	return cars, nil
}

// CarCreate method used to insert rows into car table.
func (d DB) CarCreate(ctx context.Context, car *models.Car) (*models.Car, error) {
	car.ID = uuid.New()
	car.Engine.ID = uuid.New()

	_, err := d.DB.ExecContext(ctx, create, car.ID.String(), car.Name, car.Year, car.Brand, car.Fuel, car.Engine.ID)
	if err != nil {
		return &models.Car{}, err
	}

	return car, nil
}

// CarDelete method used to delete rows from car table.
func (d DB) CarDelete(ctx context.Context, id uuid.UUID) error {
	_, err := d.DB.ExecContext(ctx, del, id.String())
	if err != nil {
		return err
	}

	return nil
}

// CarUpdate method used to update rows in car table.
func (d DB) CarUpdate(ctx context.Context, car *models.Car) (*models.Car, error) {
	_, err := d.DB.ExecContext(ctx, update, car.Name, car.Year, car.Brand, car.Fuel, car.ID)
	if err != nil {
		return &models.Car{}, err
	}

	return car, nil
}
