package car

import (
	"context"
	"developer.zopsmart.com/go/gofr/pkg/datastore"
	"developer.zopsmart.com/go/gofr/pkg/gofr"
	"github.com/DATA-DOG/go-sqlmock"
	"log"
	"testing"

	"cardealership/models"
	"developer.zopsmart.com/go/gofr/pkg/errors"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func NewMock() (*gofr.Context, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	ctx := gofr.NewContext(nil, nil, &gofr.Gofr{DataStore: datastore.DataStore{ORM: db}})
	ctx.Context = context.Background()

	return ctx, mock
}

// TestCarCreate is used to test the CarCreate method
func TestCarCreate(t *testing.T) {
	ctx, mock := NewMock()

	car := models.Car{Name: "bmwZ4", Year: 2000, Brand: "bmw", Fuel: "petrol"}

	mock.ExpectExec("INSERT INTO car(id,name,year,brand,fueltype,engine_id)"+
		"VALUES(?,?,?,?,?,?)").WithArgs(sqlmock.AnyArg(), car.Name, car.Year, car.Brand, car.Fuel, sqlmock.AnyArg()).
		WillReturnResult(sqlmock.NewResult(1, 1))

	mock.ExpectExec("INSERT INTO car(id,name,year,brand,fueltype,engine_id)"+
		"VALUES(?,?,?,?,?,?)").WithArgs(sqlmock.AnyArg(), car.Name, car.Year, car.Brand, car.Fuel, sqlmock.AnyArg()).
		WillReturnError(errors.Error("internal server error"))

	testcases := []struct {
		desc  string
		input *models.Car
		err   error
	}{
		{"success case", &car, nil},
		{"query error", &car, errors.Error("internal server error")},
	}

	s := New()

	for _, tc := range testcases {
		_, err := s.CarCreate(ctx, tc.input)
		assert.Equal(t, tc.err, err)
	}
}

// TestGet is used to test the get method of car
func TestGet(t *testing.T) {
	ctx, mock := NewMock()

	id1, err := uuid.NewRandom()
	if err != nil {
		t.Errorf("Id generation failed")
	}

	id2, _ := uuid.Parse("163744be-d162-441b-9386-43ab6042ec6d")

	testcases := []struct {
		desc     string
		id       uuid.UUID
		exOutput models.Car
		err      error
	}{
		{"success case", id1, models.Car{ID: id1, Name: "bmwZ4", Year: 2000, Brand: "bmw", Fuel: "petrol", Engine: models.Engine{}}, nil},
		{"invalid id", id2, models.Car{}, errors.Error("internal server error")},
	}

	rows := sqlmock.NewRows([]string{"Id", "Name", "Year", "Brand", "FuelType", "Engine_Id"})
	rows.AddRow(id1, "bmwZ4", 2000, "bmw", "petrol", uuid.Nil)

	mock.ExpectQuery("SELECT Id,Name,Year,Brand,FuelType,Engine_Id FROM Car where Id=?").
		WithArgs(id1).WillReturnRows(rows)
	mock.ExpectQuery("SELECT Id,Name,Year,Brand,FuelType,Engine_Id FROM car where Id=?").
		WithArgs(id2).WillReturnError(errors.Error("internal server error"))
	s := New()
	for _, tc := range testcases {
		_, err := s.CarGet(ctx, tc.id)
		assert.Equal(t, tc.err, err)
	}
}

func TestUpdate(t *testing.T) {
	ctx, mock := NewMock()

	s := New()

	id1 := uuid.New()
	id2, err := uuid.Parse("023744be-d162-441b-9386-43ab6002ec5d")

	if err != nil {
		return
	}

	car := models.Car{ID: id1, Name: "bmw", Year: 2000, Brand: "bmw", Fuel: "petrol", Engine: models.Engine{}}
	car1 := models.Car{ID: id2, Name: "ferrari f50", Year: 2000, Brand: "ferrari", Fuel: "petrol", Engine: models.Engine{}}

	testcases := []struct {
		desc     string
		input    *models.Car
		exOutput *models.Car
		err      error
	}{
		{"success case", &car, &car, nil},
		{"id not in db", &car1, &models.Car{}, errors.Error("internal server error")},
	}

	mock.ExpectExec("UPDATE Car SET Name=?,Year=?,Brand=?,FuelType=? where Id=?").
		WithArgs(car.Name, car.Year, car.Brand, car.Fuel, car.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	mock.ExpectExec("UPDATE Car SET Name=?,Year=?,Brand=?,FuelType=? where Id=?").
		WithArgs(car1.Name, car1.Year, car1.Brand, car1.Fuel, car1.ID).
		WillReturnError(errors.Error("internal server error"))

	for _, tc := range testcases {
		resp, err := s.CarUpdate(ctx, tc.input)
		assert.Equal(t, tc.err, err)
		assert.Equal(t, tc.exOutput, resp)
	}

}

// TestDB_Delete is used to test the delete method of car
func TestDelete(t *testing.T) {
	ctx, mock := NewMock()

	s := New()

	id1 := uuid.New()
	id2 := uuid.New()

	testcases := []struct {
		desc string
		id   uuid.UUID
		err  error
	}{
		{"success case", id1, nil},
		{"id not in db", id2, errors.Error("internal server error")},
	}

	mock.ExpectExec("delete from Car where Id=?").WithArgs(id1.String()).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectExec("delete from Car where Id=?").WithArgs(id2.String()).WillReturnError(errors.Error("internal server error"))

	for _, tc := range testcases {
		err := s.CarDelete(ctx, tc.id)
		assert.Equal(t, tc.err, err)
	}
}

// TestDB_GetByBrand is used to test the get method of car
func TestGetByBrandSuccess(t *testing.T) {

	ctx, mock := NewMock()
	s := New()
	id1 := uuid.New()
	id2 := uuid.New()
	car := []*models.Car{{ID: id1, Name: "ferrari f40", Year: 2015, Brand: "ferrari", Fuel: "diesel", Engine: models.Engine{ID: id2}}}
	rows := sqlmock.NewRows([]string{"id", "name", "year", "brand", "fuelType", "engine_id"})
	rows.AddRow(id1, "ferrari f40", 2015, "ferrari", "diesel", id1)

	testcases := []struct {
		desc     string
		brand    string
		isEngine bool
		quErr    error
		exOutput []*models.Car
		err      error
	}{
		{"success case", "ferrari", false, nil, car, nil},
	}
	for _, tc := range testcases {
		mock.ExpectQuery("select * from Car where Brand=?").WillReturnRows(rows)

		_, err := s.CarGetByBrand(ctx, tc.brand, tc.isEngine)

		assert.Equal(t, tc.err, err)
	}
}
func TestGetByBrandFailure(t *testing.T) {

	ctx, mock := NewMock()
	s := New()

	id1 := uuid.New()
	emptyRows := sqlmock.NewRows([]string{"Id", "Name", "Year", "Brand", "fuelType", "Engine_ID"})
	scanErrRow := sqlmock.NewRows([]string{"Id", "Name", "Year", "Fuel", "Engine_ID"}).AddRow(id1.String(),
		"bmw 2", 2018, "electric", uuid.New())
	testcases := []struct {
		desc       string
		brand      string
		isEngine   bool
		quErr      error
		outputRows *sqlmock.Rows
		exOutput   []*models.Car
		err        error
	}{
		{"failure", "ferrari", false, errors.Error("internal server error"),
			emptyRows, []*models.Car{}, errors.Error("internal server error")},
		{"scan error", "ferrari", false, nil,
			scanErrRow, []*models.Car{}, errors.Error("error while scanning")},
	}

	for _, tc := range testcases {
		mock.ExpectQuery("select * from Car where Brand=?").WithArgs(tc.brand).WillReturnRows(tc.outputRows).WillReturnError(tc.quErr)

		_, err := s.CarGetByBrand(ctx, tc.brand, tc.isEngine)

		assert.Equal(t, tc.err, err)
	}
}
