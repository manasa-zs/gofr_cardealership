package car

import (
	"context"
	"database/sql"
	"errors"
	"reflect"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"manasa/models"
)

// TestDB_CarCreate is used to test Create method of car in store layer.
func TestDB_CarCreate(t *testing.T) {
	ctx := context.Background()

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("%s", err)
	}

	a := New(db)

	defer db.Close()

	car := models.Car{Name: "b2", Year: 2001, Brand: "BMW", Fuel: "petrol"}
	queryError := errors.New("query error")

	mock.ExpectExec(create).WithArgs(sqlmock.AnyArg(), car.Name, car.Year, car.Brand, car.Fuel, sqlmock.AnyArg()).
		WillReturnResult(sqlmock.NewResult(1, 1))

	mock.ExpectExec(create).WithArgs(sqlmock.AnyArg(), car.Name, car.Year, car.Brand, car.Fuel, sqlmock.AnyArg()).
		WillReturnError(queryError)

	testcases := []struct {
		desc string
		car  models.Car
		err  error
	}{
		{"success case", car, nil},
		{"query error", models.Car{}, queryError},
	}

	for i, tc := range testcases {
		_, err := a.CarCreate(ctx, &car)
		if !reflect.DeepEqual(err, tc.err) {
			t.Errorf("[TEST %v] failed got %v want %v", i, err, tc.err)
		}
	}
}

// TestDB_CarGetById is used to test GetById method of car in store layer.
func TestDB_CarGetById(t *testing.T) {
	ctx := context.Background()

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("%s", err)
	}

	a := New(db)
	defer db.Close()

	id, err := uuid.NewRandom()
	if err != nil {
		t.Errorf("Id generate Failed")
	}

	car := models.Car{ID: id, Name: "b3", Year: 2005, Brand: "BMW", Fuel: "petrol", Engine: models.Engine{}}

	queryError := errors.New("query error")

	row := sqlmock.NewRows([]string{"id", "name", "year", "brand", "fueltype", "engine_id"}).AddRow(id, "b3", 2005, "BMW", "petrol", uuid.Nil)

	mock.ExpectQuery(getByID).WithArgs(id.String()).WillReturnRows(row)
	mock.ExpectQuery(getByID).WithArgs(uuid.Nil).WillReturnError(queryError)
	mock.ExpectQuery(getByID).WithArgs(uuid.Nil).WillReturnRows(sqlmock.NewRows([]string{"id", "name", "year",
		"brand", "fueltype", "engine_id"}).
		RowError(1, errors.New("entity not found")))

	testcases := []struct {
		desc   string
		input  uuid.UUID
		output *models.Car
		err    error
	}{
		{"success case", car.ID, &car, nil},
		{"query error", uuid.Nil, &models.Car{}, queryError},
		{"row error", uuid.Nil, &models.Car{}, errors.New("entity not found")},
	}

	for i, tc := range testcases {
		res, err := a.CarGetByID(ctx, tc.input)

		assert.Equal(t, tc.err, err)

		if !reflect.DeepEqual(res, tc.output) {
			t.Errorf("[Test %v] failed got %v Expected %v", i, res, tc.output)
		}
	}
}

// TestDB_CarUpdate is used to test Update method of car in store layer.
func TestDB_CarUpdate(t *testing.T) {
	ctx := context.Background()
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	a := New(db)

	if err != nil {
		t.Fatalf("%s", err)
	}
	defer db.Close()

	id := uuid.New()
	id1 := uuid.New()

	car := models.Car{ID: id, Name: "b4", Year: 2007, Brand: "BMW", Fuel: "petrol", Engine: models.Engine{}}
	car1 := models.Car{ID: id1, Name: "b4", Year: 2007, Brand: "BMW", Fuel: "petrol", Engine: models.Engine{}}

	mock.ExpectExec(update).WithArgs(car.Name, car.Year, car.Brand, car.Fuel, car.ID).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectExec(update).WithArgs(car1.Name, car1.Year, car1.Brand, car1.Fuel, car1.ID).WillReturnError(errors.New("scan error"))

	testcases := []struct {
		desc   string
		input  *models.Car
		output *models.Car
		err    error
	}{
		{"success case", &car, &car, nil},
		{"scan error", &car1, &models.Car{}, errors.New("scan error")},
	}

	for i, tc := range testcases {
		res, err := a.CarUpdate(ctx, tc.input)

		if !reflect.DeepEqual(res, tc.output) {
			t.Errorf("[TEST %v] failed got %v want %v", i, res, tc.output)
		}

		if !reflect.DeepEqual(err, tc.err) {
			t.Errorf("[TEST %v] failed got %v want %v", i, err, tc.err)
		}
	}
}

// TestDB_CarDelete is used to test Delete method of car in store layer.
func TestDB_CarDelete(t *testing.T) {
	ctx := context.Background()

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("%s", err)
	}

	a := New(db)

	defer db.Close()

	id := uuid.New()
	id1 := uuid.New()

	mock.ExpectExec(del).WithArgs(id.String()).WillReturnResult(sqlmock.NewResult(
		1, 1))

	mock.ExpectExec(del).WithArgs(id1.String()).WillReturnError(errors.New("id doesn't exists"))

	cases := []struct {
		desc string
		id   uuid.UUID
		err  error
	}{
		{"success case", id, nil},
		{"invalid error", id1, errors.New("id doesn't exists")},
	}

	for i, tc := range cases {
		err := a.CarDelete(ctx, tc.id)
		if !reflect.DeepEqual(err, tc.err) {
			t.Errorf("\n[TEST %v] Failed Got %v\n Expected %v", i, err, tc.err)
		}
	}
}

// TestDB_GetByBrand is used to test GetByBrand method of car in store layer.
func TestDB_CarGetByBrand(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Errorf("%s error occurred while connecting to database", err)
	}

	id1 := uuid.New()
	engineID := uuid.New()
	id2 := uuid.New()
	engineID2 := uuid.New()

	result := []*models.Car{
		{ID: id1, Name: "Model X", Year: 2018, Brand: "Tesla", Fuel: "Electric", Engine: models.Engine{ID: engineID}},
		{ID: id2, Name: "Model Y", Year: 2021, Brand: "Tesla", Fuel: "Petrol", Engine: models.Engine{ID: engineID2}}}

	rows := sqlmock.NewRows([]string{"id", "name", "year", "brand", "fueltype", "engineID"}).AddRow(id1.String(),
		"Model X", "2018", "Tesla", "Electric", engineID.String()).AddRow(id2.String(), "Model Y", "2021", "Tesla",
		"Petrol", engineID2.String())

	emptyRows := sqlmock.NewRows([]string{"id", "name", "year", "brand", "fueltype", "engineID"})
	scanErrRow := sqlmock.NewRows([]string{"id", "name", "year", "fueltype", "engineID"}).AddRow(id1.String(),
		"Model X", "2018", "Electric", engineID.String())

	testCases := []struct {
		desc       string
		input      string
		outputRows *sqlmock.Rows
		output     []*models.Car
		rowErr     error

		err error
	}{
		{"Success", "Tesla", rows, result, nil, nil},
		{"Error", "Tesla", emptyRows, nil, errors.New(
			"query error"), errors.New("query error")},
		{"Error", "Tesla", scanErrRow, nil, nil, errors.New(
			"error while scanning")},
	}
	carStore := New(db)

	for i, tc := range testCases {
		mock.ExpectQuery("select * from car where brand = ?").WithArgs(tc.input).WillReturnRows(
			tc.outputRows).WillReturnError(tc.rowErr)

		result, err := carStore.CarGetByBrand(context.TODO(), tc.input)

		if !reflect.DeepEqual(err, tc.err) {
			t.Errorf("Testcase[%v] failed desc: %v\n Expected error : %v Output: %v\n ", i, tc.desc, tc.err, err)
		}

		assert.Equal(t, tc.output, result)
	}
}

func TestDB_CarGetByBrandRowError(t *testing.T) {
	ctx := context.Background()

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("%s", err)
	}

	a := New(db)
	defer db.Close()

	mock.ExpectQuery(getByBrand).WithArgs("BMW").WillReturnError(sql.ErrNoRows)

	_, err = a.CarGetByBrand(ctx, "BMW")

	assert.Equal(t, errors.New("entity not found"), err)
}
