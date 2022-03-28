package engine

import (
	"context"
	"errors"
	"fmt"
	"log"
	"reflect"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"manasa/models"
)

// TestDB_EngineCreate is used to test Create method of engine in store layer.
func TestDBEngineCreate(t *testing.T) { //nolint
	ctx := context.Background()

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		log.Fatalf(err.Error())
	}

	a := New(db)

	defer db.Close()

	id, err := uuid.NewUUID()
	if err != nil {
		fmt.Println(err)
	}

	ID, err := uuid.NewUUID()
	if err != nil {
		fmt.Println(err)
	}

	testcases := []struct {
		desc   string
		input  *models.Engine
		output *models.Engine
		err    error
	}{
		{"success case", &models.Engine{ID: id, Displacement: 400, NoOfCylinders: 0, Range: 0},
			&models.Engine{ID: id, Displacement: 400, NoOfCylinders: 0, Range: 0}, nil},
		{"query error", &models.Engine{ID: ID, Displacement: 400, NoOfCylinders: 0, Range: 0},
			&models.Engine{}, err},
	}

	for i, tc := range testcases {
		if tc.input.ID == id {
			mock.ExpectExec(create).
				WithArgs(tc.output.ID, tc.output.Displacement, tc.output.NoOfCylinders, tc.output.Range).
				WillReturnResult(sqlmock.NewResult(1, 1))
		} else {
			mock.ExpectExec(create).
				WithArgs(tc.output.ID, tc.output.Displacement, tc.output.NoOfCylinders, tc.output.Range).
				WillReturnError(tc.err)
		}

		resp, err := a.EngineCreate(ctx, tc.input)

		if tc.err != nil {
			if !reflect.DeepEqual(err.Error(), tc.err.Error()) {
				t.Errorf("[TEST %v]: got %v want %v", i, err.Error(), tc.err.Error())
			}
		} else {
			if !reflect.DeepEqual(resp, tc.output) {
				t.Errorf("[TEST %v]: got %v want %v", i, resp, tc.output)
			}
		}
	}
}

// TestDB_EngineGetByID is used to test GetById method of engine in store layer.
func TestDBEngineGetByID(t *testing.T) {
	ctx := context.Background()

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		log.Fatalln(err)
	}

	a := New(db)
	defer db.Close()

	id := uuid.New()
	ID := uuid.New()

	row := sqlmock.NewRows([]string{"engine_id", "displacement", "no_of_cylinders", "range"}).AddRow(uuid.Nil, 300, 2, 0)

	testcases := []struct {
		desc   string
		id     uuid.UUID
		output *models.Engine
		err    error
	}{
		{"success case", id, &models.Engine{
			Displacement:  300,
			Range:         0,
			NoOfCylinders: 2,
		}, nil},
		{"error", ID, &models.Engine{}, errors.New("internal server error")},
	}
	for _, tc := range testcases {
		if tc.id == id {
			mock.ExpectQuery(getByID).WithArgs(id.String()).WillReturnRows(row)
		} else {
			mock.ExpectQuery(getByID).WithArgs(ID.String()).WillReturnError(tc.err)
		}

		resp, err := a.EngineGetByID(ctx, tc.id)

		assert.Equal(t, err, tc.err)
		assert.Equal(t, resp, tc.output)
	}
}

// TestDB_EngineDelete is used to test Delete method of engine in store layer.
func TestDB_EngineDelete(t *testing.T) {
	ctx := context.Background()

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		log.Fatalln(err)
	}

	a := New(db)
	defer db.Close()

	id, err := uuid.NewRandom()
	if err != nil {
		t.Errorf("failed to generate id %v", err)
		return
	}

	deleteErr := errors.New("delete failed")

	mock.ExpectExec(del).WithArgs(id.String()).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectExec(del).WithArgs(uuid.Nil).WillReturnError(deleteErr)

	testcases := []struct {
		desc string
		id   uuid.UUID
		err  error
	}{
		{"success case", id, nil},
		{"delete error", uuid.Nil, deleteErr},
	}

	for i, tc := range testcases {
		err := a.EngineDelete(ctx, tc.id)
		if err != tc.err {
			t.Errorf("[Test %v] Got %v Expected %v", i, err, tc.err)
		}
	}
}

// TestDB_EngineUpdate is used to test Delete method of engine in store layer.
func TestEngineStore_EngineUpdate(t *testing.T) {
	ctx := context.Background()

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		log.Fatalln(err)
	}

	a := New(db)
	defer db.Close()

	id := uuid.New()
	ID := uuid.New()

	engine := models.Engine{ID: id, Displacement: 300, NoOfCylinders: 4, Range: 0}
	engine1 := models.Engine{ID: ID, Displacement: 300, NoOfCylinders: 4, Range: 0}

	mock.ExpectExec(update).WithArgs(engine.Displacement, engine.NoOfCylinders, engine.Range, engine.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectExec(update).WithArgs(engine1.Displacement, engine1.NoOfCylinders, engine1.Range, engine1.ID).
		WillReturnError(errors.New("scan error"))

	testcases := []struct {
		desc  string
		input *models.Engine
		err   error
	}{
		{"success case", &engine, nil},
		{"scan error", &engine1, errors.New("scan error")},
	}
	for i, tc := range testcases {
		_, err := a.EngineUpdate(ctx, tc.input)
		if !reflect.DeepEqual(err, tc.err) {
			t.Errorf("[Test %v]Failed Got %v Expected %v", i, err, tc.err)
		}
	}
}
