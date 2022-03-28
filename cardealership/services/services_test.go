package services

import (
	"cardealership/models"
	"developer.zopsmart.com/go/gofr/pkg/errors"
	"developer.zopsmart.com/go/gofr/pkg/gofr"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

// TestValidation is used to test the validate function
func TestValidation(t *testing.T) {
	id := uuid.New()
	testcases := []struct {
		desc  string
		input models.Car
		err   error
	}{
		{desc: "success case", input: models.Car{ID: id, Name: "bmwZ4", Year: 2000, Brand: "bmw", Fuel: "petrol",
			Engine: models.Engine{ID: id, Displacement: 20, NoOfCylinders: 5, Range: 0}},
			err: nil},
		{desc: "year not in the given limit",
			input: models.Car{ID: id, Name: "bmwZ4", Year: 1970, Brand: "bmw", Fuel: "petrol",
				Engine: models.Engine{ID: id, Displacement: 20, NoOfCylinders: 5, Range: 0}},
			err: errors.Error("year should be between 1980 and 2022")},
		{desc: "invalid brand", input: models.Car{ID: id, Name: "tataZ4", Year: 2000, Brand: "tata", Fuel: "petrol",
			Engine: models.Engine{ID: id, Displacement: 20, NoOfCylinders: 5, Range: 0}},
			err: errors.Error("invalid brand")},
		{desc: "invalid fuel type", input: models.Car{ID: id, Name: "bmwZ4", Year: 2000, Brand: "bmw", Fuel: "oil",
			Engine: models.Engine{ID: id, Displacement: 20, NoOfCylinders: 5, Range: 0}},
			err: errors.Error("invalid fuel")},
		{desc: "displacement cannot be negative", input: models.Car{ID: id, Name: "bmwZ4", Year: 2000, Brand: "bmw", Fuel: "petrol",
			Engine: models.Engine{ID: id, Displacement: -2, NoOfCylinders: 5, Range: 0}},
			err: errors.Error("displacement should be positive")},
		{desc: "noOfCylinders cannot be negative", input: models.Car{ID: id, Name: "bmwZ4", Year: 2000, Brand: "bmw", Fuel: "petrol",
			Engine: models.Engine{ID: id, Displacement: 2, NoOfCylinders: -5, Range: 0}},
			err: errors.Error("noOfCylinders should be positive")},
		{desc: "range cannot be negative", input: models.Car{ID: id, Name: "bmwZ4", Year: 2000, Brand: "bmw", Fuel: "electric",
			Engine: models.Engine{ID: id, Displacement: 0, NoOfCylinders: 0, Range: -1}},
			err: errors.Error("range should be positive")},
		{desc: "range cannot be for petrol type", input: models.Car{ID: id, Name: "bmwZ4", Year: 2000, Brand: "bmw", Fuel: "petrol",
			Engine: models.Engine{ID: id, Displacement: 10, NoOfCylinders: 2, Range: 3}},
			err: errors.Error("given fuel type does not support range")},
		{desc: "electric cannot have displacement", input: models.Car{ID: id, Name: "bmwZ4", Year: 2000, Brand: "bmw", Fuel: "electric",
			Engine: models.Engine{ID: id, Displacement: 10, NoOfCylinders: 0, Range: 1}},
			err: errors.Error("electric cannot have displacement")},
		{desc: "electric cannot have  noOfCylinders", input: models.Car{ID: id, Name: "bmwZ4", Year: 2000, Brand: "bmw", Fuel: "electric",
			Engine: models.Engine{ID: id, Displacement: 0, NoOfCylinders: 10, Range: 1}},
			err: errors.Error("electric cannot have  noOfCylinders")},
	}

	for _, tc := range testcases {
		err := Validation(tc.input)
		assert.Equal(t, err, tc.err)
	}
}

// TestCreateValidation is used to validate the input of create method
func TestCreateValidation(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockCar := NewMockCar(ctrl)
	mockEngine := NewMockEngine(ctrl)
	svc := New(mockCar, mockEngine)
	id := uuid.New()
	id1 := uuid.New()
	testcases := []struct {
		desc  string
		input models.Car
		err   error
	}{
		{desc: "invalid brand", input: models.Car{ID: id, Name: "tataZ4", Year: 2000, Brand: "tata", Fuel: "petrol",
			Engine: models.Engine{ID: id1, Displacement: 20, NoOfCylinders: 5, Range: 0}},
			err: errors.Error("invalid brand")},
	}

	for _, tc := range testcases {
		ctx := gofr.NewContext(nil, nil, gofr.New())
		_, err := svc.Create(ctx, &tc.input)
		assert.Equal(t, err, tc.err)
	}
}

//TestCreateSuccess is used to test the Success case of Create
func TestCreateSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	id := uuid.New()
	id1 := uuid.New()
	testcases := []struct {
		desc     string
		input    models.Car
		carError error
		enError  error
		err      error
	}{
		{desc: "success case", input: models.Car{ID: id, Name: "bmwZ4", Year: 2000, Brand: "bmw", Fuel: "petrol",
			Engine: models.Engine{ID: id1, Displacement: 20, NoOfCylinders: 5, Range: 0}}, carError: nil, enError: nil, err: nil},
	}

	ct := gomock.NewController(t)
	mockCar := NewMockCar(ct)
	mockEngine := NewMockEngine(ct)
	svc := New(mockCar, mockEngine)

	for _, tc := range testcases {
		ctx := gofr.NewContext(nil, nil, gofr.New())
		mockCar.EXPECT().CarCreate(ctx, &tc.input).Return(&tc.input, tc.carError)
		mockEngine.EXPECT().EngineCreate(ctx, &tc.input.Engine).Return(&tc.input.Engine, tc.enError)
		_, err := svc.Create(ctx, &tc.input)
		assert.Equal(t, err, tc.err)
	}
}

// TestCreateCarError is used to test the Create method with error in store layer
func TestCreateCarError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	id := uuid.New()
	id1 := uuid.New()
	testcases := []struct {
		desc     string
		input    models.Car
		carError error
		enError  error
		err      error
	}{

		{desc: "error in car store layer", input: models.Car{ID: id, Name: "bmwZ4", Year: 2000, Brand: "bmw", Fuel: "petrol",
			Engine: models.Engine{ID: id1, Displacement: 20, NoOfCylinders: 5, Range: 0}},
			carError: errors.Error("error in car layer"), enError: nil, err: errors.Error("error in car layer")},
	}

	ct := gomock.NewController(t)
	mockCar := NewMockCar(ct)
	mockEngine := NewMockEngine(ct)
	svc := New(mockCar, mockEngine)

	for _, tc := range testcases {
		ctx := gofr.NewContext(nil, nil, gofr.New())
		mockCar.EXPECT().CarCreate(ctx, &tc.input).Return(&tc.input, tc.carError)
		_, err := svc.Create(ctx, &tc.input)
		assert.Equal(t, err, tc.err)
	}
}

// TestCreateEngineError is used to test the create method with error in store layer
func TestCreateEngineError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	id := uuid.New()
	id1 := uuid.New()
	testcases := []struct {
		desc     string
		input    models.Car
		carError error
		enError  error
		err      error
	}{

		{desc: "error in engine store layer", input: models.Car{ID: id, Name: "bmwZ4", Year: 2000, Brand: "bmw", Fuel: "petrol",
			Engine: models.Engine{ID: id1, Displacement: 20, NoOfCylinders: 5, Range: 0}},
			carError: nil, enError: errors.Error("error in engine layer"), err: errors.Error("error in engine layer")},
	}

	ct := gomock.NewController(t)
	mockCar := NewMockCar(ct)
	mockEngine := NewMockEngine(ct)
	svc := New(mockCar, mockEngine)

	for _, tc := range testcases {
		ctx := gofr.NewContext(nil, nil, gofr.New())
		mockCar.EXPECT().CarCreate(ctx, &tc.input).Return(&tc.input, tc.carError)
		mockEngine.EXPECT().EngineCreate(ctx, &tc.input.Engine).Return(&tc.input.Engine, tc.enError)
		_, err := svc.Create(ctx, &tc.input)
		assert.Equal(t, err, tc.err)
	}
}

// TestGetByIdSuccess is used to test the success case of GetByID
func TestGetByIdSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	id1, _ := uuid.Parse("063744be-d162-441b-9386-43ab6042ed6f")
	id2, _ := uuid.Parse("063744be-d162-441b-9386-43ab6042ec5d")

	testcases := []struct {
		desc     string
		id       uuid.UUID
		exOutput models.Car
		err      error
	}{
		{"success case", id1, models.Car{ID: id1, Name: "bmwZ4", Year: 2000, Brand: "bmw", Fuel: "petrol",
			Engine: models.Engine{ID: id2, Displacement: 20, NoOfCylinders: 5}}, nil},
	}
	ct := gomock.NewController(t)
	mockCar := NewMockCar(ct)
	mockEngine := NewMockEngine(ct)
	svc := New(mockCar, mockEngine)

	for _, tc := range testcases {
		ctx := gofr.NewContext(nil, nil, gofr.New())
		mockEngine.EXPECT().EngineGet(ctx, tc.exOutput.Engine.ID).Return(&tc.exOutput.Engine, tc.err)
		mockCar.EXPECT().CarGet(ctx, tc.id).Return(&tc.exOutput, tc.err)
		resp, _ := svc.GetByID(ctx, tc.id)
		assert.Equal(t, resp, &tc.exOutput)
	}
}

// TestGetByID is used to check getById method
func TestGetByNilID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	testcases := []struct {
		desc     string
		id       uuid.UUID
		exOutput models.Car
		err      error
	}{

		{"invalid id", uuid.Nil, models.Car{}, errors.Error("missing id")},
	}

	ct := gomock.NewController(t)
	mockCar := NewMockCar(ct)
	mockEngine := NewMockEngine(ct)
	svc := New(mockCar, mockEngine)

	for _, tc := range testcases {
		ctx := gofr.NewContext(nil, nil, gofr.New())
		mockCar.EXPECT().CarGet(ctx, tc.id).Return(&tc.exOutput, tc.err)
		resp, _ := svc.GetByID(ctx, tc.id)
		assert.Equal(t, resp, &tc.exOutput)
	}
}

// TestGetByID is used to check getById method
func TestGetByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	id, _ := uuid.NewRandom()

	testcases := []struct {
		desc     string
		id       uuid.UUID
		exOutput models.Car
		err1     error
		err2     error
	}{
		{"engine id doesn't exist", id, models.Car{}, nil, errors.Error("invalid engine id")},
	}

	ct := gomock.NewController(t)
	mockCar := NewMockCar(ct)
	mockEngine := NewMockEngine(ct)
	svc := New(mockCar, mockEngine)

	for _, tc := range testcases {
		ctx := gofr.NewContext(nil, nil, gofr.New())
		mockCar.EXPECT().CarGet(ctx, tc.id).Return(&tc.exOutput, tc.err1)
		mockEngine.EXPECT().EngineGet(ctx, uuid.Nil).Return(&tc.exOutput.Engine, tc.err2)
		resp, _ := svc.GetByID(ctx, tc.id)
		assert.Equal(t, resp, &tc.exOutput)
	}
}

// TestDelete is used to test the delete method
func TestDelete(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	id := uuid.New()
	testcases := []struct {
		desc   string
		id     uuid.UUID
		idErr  error
		CarErr error
		EngErr error
		err    error
	}{

		{"error from datastore layer", id, errors.Error("error from datastore layer"),
			nil, nil, errors.Error("error from datastore layer")},
	}
	car := models.Car{ID: id, Name: "bmw", Year: 2000, Brand: "bmw", Fuel: "electric",
		Engine: models.Engine{Displacement: 0, NoOfCylinders: 0, Range: 10}}

	mockCar := NewMockCar(ctrl)
	mockEngine := NewMockEngine(ctrl)
	svc := New(mockCar, mockEngine)

	for i, tc := range testcases {
		ctx := gofr.NewContext(nil, nil, gofr.New())
		mockCar.EXPECT().CarGet(ctx, tc.id).Return(&car, tc.idErr)
		err := svc.Delete(ctx, tc.id)

		if !reflect.DeepEqual(err, tc.err) {
			t.Errorf("[Test %v] Failed. Got %v Expected %v", i+1, err, tc.err)
		}
	}
}

// TestDeleteSuccess is used to test the delete method
func TestDeleteSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	id := uuid.New()
	testcases := []struct {
		desc   string
		id     uuid.UUID
		idErr  error
		CarErr error
		EngErr error
		err    error
	}{
		{"success case", id, nil, nil, nil, nil},
		{"invalid id", id, nil, errors.Error("invalid id"), nil, errors.Error("invalid id")},
	}
	car := models.Car{ID: id, Name: "bmw", Year: 2000, Brand: "bmw", Fuel: "electric",
		Engine: models.Engine{Displacement: 0, NoOfCylinders: 0, Range: 10}}

	mockCar := NewMockCar(ctrl)
	mockEngine := NewMockEngine(ctrl)
	svc := New(mockCar, mockEngine)

	for i, tc := range testcases {
		ctx := gofr.NewContext(nil, nil, gofr.New())
		mockCar.EXPECT().CarGet(ctx, tc.id).Return(&car, tc.idErr)
		mockEngine.EXPECT().EngineDelete(ctx, car.Engine.ID).Return(tc.EngErr)
		mockCar.EXPECT().CarDelete(ctx, tc.id).Return(tc.CarErr)

		err := svc.Delete(ctx, tc.id)

		if !reflect.DeepEqual(err, tc.err) {
			t.Errorf("[Test %v] Failed. Got %v Expected %v", i+1, err, tc.err)
		}
	}
}

// TestDeleteEngError is used to test the delete method
func TestDeleteEngError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	id := uuid.New()
	testcases := []struct {
		desc   string
		id     uuid.UUID
		idErr  error
		CarErr error
		EngErr error
		err    error
	}{
		{"invalid id", id, nil, nil, errors.Error("invalid id"), errors.Error("invalid id")},
	}
	car := models.Car{ID: id, Name: "bmw", Year: 2000, Brand: "bmw", Fuel: "electric",
		Engine: models.Engine{Displacement: 0, NoOfCylinders: 0, Range: 10}}

	mockCar := NewMockCar(ctrl)
	mockEngine := NewMockEngine(ctrl)
	svc := New(mockCar, mockEngine)

	for i, tc := range testcases {
		ctx := gofr.NewContext(nil, nil, gofr.New())
		mockCar.EXPECT().CarGet(ctx, tc.id).Return(&car, tc.idErr)
		mockEngine.EXPECT().EngineDelete(ctx, car.Engine.ID).Return(tc.EngErr)

		err := svc.Delete(ctx, tc.id)

		if !reflect.DeepEqual(err, tc.err) {
			t.Errorf("[Test %v] Failed. Got %v Expected %v", i+1, err, tc.err)
		}
	}
}

// TestUpdateValidation is used to test the update method
func TestUpdateValidation(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	id := uuid.New()

	testcases := []struct {
		desc     string
		input    models.Car
		exOutput models.Car
		err      error
	}{

		{desc: "year not in the given limit", input: models.Car{ID: id, Name: "bmwZ4", Year: 1970, Brand: "bmw", Fuel: "petrol",
			Engine: models.Engine{ID: id, Displacement: 20, NoOfCylinders: 5, Range: 0}},
			exOutput: models.Car{}, err: errors.Error("year should be between 1980 and 2022")},
	}

	ct := gomock.NewController(t)
	mockCar := NewMockCar(ct)
	mockEngine := NewMockEngine(ct)
	svc := New(mockCar, mockEngine)

	for _, tc := range testcases {
		ctx := gofr.NewContext(nil, nil, gofr.New())
		resp, err := svc.Update(ctx, &tc.input)
		assert.Equal(t, resp, &tc.exOutput)
		assert.Equal(t, err, tc.err)
	}
}

// TestUpdate is used to test Update method
func TestUpdate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	id1, _ := uuid.Parse("063744be-d162-441b-9386-43ab6042ec5d")
	car := models.Car{ID: id1, Name: "bmw", Year: 2000, Brand: "bmw", Fuel: "petrol",
		Engine: models.Engine{ID: id1, Displacement: 20, NoOfCylinders: 5, Range: 0}}

	testcases := []struct {
		desc     string
		input    models.Car
		exOutput models.Car

		err error
	}{
		{"success case", car, car, nil},
	}

	ct := gomock.NewController(t)
	mockCar := NewMockCar(ct)
	mockEngine := NewMockEngine(ct)
	svc := New(mockCar, mockEngine)

	for _, tc := range testcases {
		ctx := gofr.NewContext(nil, nil, gofr.New())
		mockCar.EXPECT().CarGet(ctx, tc.input.ID).Return(&tc.exOutput, tc.err)
		mockEngine.EXPECT().EngineUpdate(ctx, &tc.input.Engine).Return(&tc.exOutput.Engine, tc.err)
		mockCar.EXPECT().CarUpdate(ctx, &tc.input).Return(&tc.exOutput, tc.err)

		resp, err := svc.Update(ctx, &tc.input)
		assert.Equal(t, resp, &tc.exOutput)
		assert.Equal(t, err, tc.err)
	}
}

// TestCarUpdateError is used to test the Update method in service layer
func TestCarUpdateError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	id := uuid.New()
	id1 := uuid.New()
	car := models.Car{ID: id, Name: "bmw", Year: 2000, Brand: "bmw", Fuel: "petrol",
		Engine: models.Engine{ID: id1, Displacement: 20, NoOfCylinders: 5, Range: 0}}

	testcases := []struct {
		desc     string
		input    models.Car
		exOutput models.Car
		err      error
	}{
		{"error in store layer", car, models.Car{}, errors.Error("error in store")},
	}

	ct := gomock.NewController(t)
	mockCar := NewMockCar(ct)
	mockEngine := NewMockEngine(ct)
	svc := New(mockCar, mockEngine)

	for _, tc := range testcases {
		ctx := gofr.NewContext(nil, nil, gofr.New())
		mockCar.EXPECT().CarUpdate(ctx, &tc.input).Return(&tc.exOutput, tc.err)
		resp, err := svc.Update(ctx, &tc.input)
		assert.Equal(t, resp, &tc.exOutput)
		assert.Equal(t, err, tc.err)
	}
}

// TestCarUpdateInvalidID is used to test the update method
func TestCarUpdateInvalidID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	id := uuid.New()
	id1 := uuid.New()
	car := models.Car{ID: id, Name: "bmw", Year: 2000, Brand: "bmw", Fuel: "petrol",
		Engine: models.Engine{ID: id1, Displacement: 20, NoOfCylinders: 5, Range: 0}}

	testcases := []struct {
		desc     string
		input    models.Car
		exOutput models.Car
		carErr   error
		err      error
	}{
		{"error in id", car, models.Car{}, nil, errors.Error("error in id")},
	}

	ct := gomock.NewController(t)
	mockCar := NewMockCar(ct)
	mockEngine := NewMockEngine(ct)
	svc := New(mockCar, mockEngine)

	for _, tc := range testcases {
		ctx := gofr.NewContext(nil, nil, gofr.New())
		mockCar.EXPECT().CarUpdate(ctx, &tc.input).Return(&tc.exOutput, tc.carErr)
		mockCar.EXPECT().CarGet(ctx, tc.input.ID).Return(&tc.exOutput, tc.err)

		resp, err := svc.Update(ctx, &tc.input)
		assert.Equal(t, resp, &tc.exOutput)
		assert.Equal(t, err, tc.err)
	}
}

//  TestUpdateEngine is used to test the update method which has  error in engine store layer
func TestUpdateEngineErr(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	id1 := uuid.New()
	car := models.Car{ID: id1, Name: "bmw", Year: 2000, Brand: "bmw", Fuel: "petrol",
		Engine: models.Engine{ID: id1, Displacement: 20, NoOfCylinders: 5, Range: 0}}

	testcases := []struct {
		desc  string
		input models.Car

		exOutput models.Car
		carErr   error
		err      error
	}{
		{"error in engine store", car, models.Car{}, nil, errors.Error("error in engine")},
	}

	mockCar := NewMockCar(ctrl)
	mockEngine := NewMockEngine(ctrl)
	svc := New(mockCar, mockEngine)

	for _, tc := range testcases {
		ctx := gofr.NewContext(nil, nil, gofr.New())
		mockCar.EXPECT().CarUpdate(ctx, &tc.input).Return(&tc.exOutput, tc.carErr)
		mockCar.EXPECT().CarGet(ctx, tc.input.ID).Return(&tc.exOutput, tc.carErr)
		mockEngine.EXPECT().EngineUpdate(ctx, &tc.input.Engine).Return(&tc.exOutput.Engine, tc.err)

		resp, err := svc.Update(ctx, &tc.input)
		assert.Equal(t, resp, &tc.exOutput)
		assert.Equal(t, err, tc.err)
	}
}

// TestGet is used to test Get method
func TestGet(t *testing.T) { //nolint
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	id1, _ := uuid.Parse("063744be-d162-441b-9386-43ab6042ec5d")

	testcases := []struct {
		desc     string
		id       uuid.UUID
		brand    string
		engine   bool
		exOutput []*models.Car
		err      error
	}{
		{"error in store layer", id1, "ferrari", false, []*models.Car{}, errors.Error("error from database layer")},
		{"invalid brand", id1, "tata", false, []*models.Car{}, errors.Error("invalid brand")},
	}

	mockCar := NewMockCar(ctrl)
	mockEngine := NewMockEngine(ctrl)
	svc := New(mockCar, mockEngine)

	for i, tc := range testcases {
		ctx := gofr.NewContext(nil, nil, gofr.New())
		if tc.brand == "bmw" || tc.brand == "tesla" || tc.brand == "ferrari" || tc.brand == "porsche" || tc.brand == "mercedes" {
			mockCar.EXPECT().CarGetByBrand(ctx, tc.brand, tc.engine).Return(tc.exOutput, tc.err)

			if tc.engine == true {
				mockEngine.EXPECT().EngineGet(ctx, tc.id).Return(&tc.exOutput[i].Engine, tc.err)
			}
		}

		resp, err := svc.GetByBrand(ctx, tc.brand, tc.engine)
		assert.Equal(t, &resp, &tc.exOutput)
		assert.Equal(t, err, tc.err)
	}
}

// TestGetEngineError is used to test the GetByBrand with error in engine get
func TestGetEngineError(t *testing.T) { //nolint
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	id1 := uuid.New()

	car2 := []*models.Car{{ID: id1, Name: "ferrari f40", Year: 2015, Brand: "ferrari", Fuel: "diesel",
		Engine: models.Engine{ID: id1, Displacement: 20, NoOfCylinders: 2, Range: 0}}}

	testcases := []struct {
		desc     string
		id       uuid.UUID
		brand    string
		engine   bool
		exOutput []*models.Car
		err      error
	}{

		{"error in engine get", id1, "tesla", true, car2, errors.Error("error in engine get")},
	}

	mockCar := NewMockCar(ctrl)
	mockEngine := NewMockEngine(ctrl)
	svc := New(mockCar, mockEngine)

	for _, tc := range testcases {
		ctx := gofr.NewContext(nil, nil, gofr.New())
		mockCar.EXPECT().CarGetByBrand(ctx, tc.brand, tc.engine).Return(tc.exOutput, nil)
		mockEngine.EXPECT().EngineGet(ctx, tc.id).Return(nil, tc.err)
		_, err := svc.GetByBrand(ctx, tc.brand, tc.engine)
		assert.Equal(t, err, tc.err)
	}
}
