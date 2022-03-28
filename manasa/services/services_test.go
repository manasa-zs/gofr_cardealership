package services

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"manasa/models"
)

// TestService_Create method is used for testing Update method in Service layer.
func TestService_Create(t *testing.T) {
	ctx := context.Background()
	ct := gomock.NewController(t)
	mockCar := NewMockCar(ct)
	mockEngine := NewMockEngine(ct)
	svc := New(mockCar, mockEngine)

	id := uuid.New()
	car := models.Car{ID: id, Name: "b3", Year: 2001, Brand: "bmw", Fuel: "petrol", Engine: models.Engine{ID: id, Displacement: 500,
		NoOfCylinders: 200, Range: 0}}

	mockCar.EXPECT().CarCreate(ctx, &car).Return(&car, nil)
	mockEngine.EXPECT().EngineCreate(ctx, &car.Engine).Return(&car.Engine, nil)

	res, err := svc.Create(ctx, &car)

	assert.Equal(t, &car, res)
	assert.Equal(t, nil, err)
}

// TestService_CreateInvalidCar method is used for testing Update method in Service layer.
func TestService_CreateInvalidBrand(t *testing.T) {
	ctx := context.Background()
	ct := gomock.NewController(t)

	id := uuid.New()
	car := models.Car{ID: id, Name: "b3", Year: 2001, Brand: "maruthi", Fuel: "petrol", Engine: models.Engine{ID: id, Displacement: 500,
		NoOfCylinders: 200, Range: 0}}

	mockCar := NewMockCar(ct)
	mockEngine := NewMockEngine(ct)
	svc := New(mockCar, mockEngine)
	res, err := svc.Create(ctx, &car)

	assert.Equal(t, &models.Car{}, res)
	assert.Equal(t, errors.New("invalid brand"), err)
}

// TestService_CreateCarError method is used for testing Update method in Service layer.
func TestService_CreateCarError(t *testing.T) {
	ctx := context.Background()
	ct := gomock.NewController(t)

	mockCar := NewMockCar(ct)
	mockEngine := NewMockEngine(ct)
	svc := New(mockCar, mockEngine)

	id := uuid.New()

	car := models.Car{ID: id, Name: "b3", Year: 2001, Brand: "bmw", Fuel: "petrol", Engine: models.Engine{ID: id, Displacement: 500,
		NoOfCylinders: 200, Range: 0}}

	mockCar.EXPECT().CarCreate(ctx, &car).Return(&models.Car{}, errors.New("error from car in datastore layer"))

	res, err := svc.Create(ctx, &car)

	assert.Equal(t, &models.Car{}, res)
	assert.Equal(t, errors.New("error from car in datastore layer"), err)
}

// TestService_CreateEngineError method is used for testing Update method in Service layer.
func TestService_CreateEngineError(t *testing.T) {
	ctx := context.Background()
	ct := gomock.NewController(t)
	mockCar := NewMockCar(ct)
	mockEngine := NewMockEngine(ct)
	svc := New(mockCar, mockEngine)

	id := uuid.New()
	car := models.Car{ID: id, Name: "b3", Year: 2001, Brand: "bmw", Fuel: "petrol", Engine: models.Engine{ID: id, Displacement: 500,
		NoOfCylinders: 200, Range: 0}}

	mockCar.EXPECT().CarCreate(ctx, &car).Return(&car, nil)
	mockEngine.EXPECT().EngineCreate(ctx, &car.Engine).Return(&models.Engine{}, errors.New("error from datastore layer"))

	res, err := svc.Create(ctx, &car)

	assert.Equal(t, &models.Car{}, res)
	assert.Equal(t, errors.New("error from datastore layer"), err)
}

// TestService_GetByBrand method is used for testing GetByBrand method in Service layer.
func TestService_GetByBrand(t *testing.T) {
	ctx := context.Background()
	ct := gomock.NewController(t)
	mockCar := NewMockCar(ct)
	mockEngine := NewMockEngine(ct)
	svc := New(mockCar, mockEngine)

	id := uuid.New()
	car := []*models.Car{{ID: id, Name: "b3", Year: 2000, Brand: "bmw", Fuel: "petrol", Engine: models.Engine{ID: id, Displacement: 200,
		NoOfCylinders: 3, Range: 0}}}
	isEngine := true

	mockCar.EXPECT().CarGetByBrand(ctx, "bmw").Return(car, nil)
	mockEngine.EXPECT().EngineGetByID(ctx, id).Return(&car[0].Engine, nil)

	res, err := svc.GetByBrand(ctx, "bmw", isEngine)
	assert.Equal(t, nil, err)
	assert.Equal(t, car, res)
}

// TestService_GetByBrandError method is used for testing GetByBrand method in Service layer.
func TestService_GetByBrandError(t *testing.T) {
	ctx := context.Background()
	ct := gomock.NewController(t)
	mockCar := NewMockCar(ct)
	mockEngine := NewMockEngine(ct)
	svc := New(mockCar, mockEngine)

	isEngine := true

	res, err := svc.GetByBrand(ctx, "maruthi", isEngine)
	assert.Equal(t, errors.New("invalid brand"), err)
	assert.Equal(t, []*models.Car{}, res)
}

// TestService_GetByBrandCarError method is used for testing GetByBrand method in Service layer.
func TestService_GetByBrandCarError(t *testing.T) {
	ctx := context.Background()
	ct := gomock.NewController(t)
	mockCar := NewMockCar(ct)
	mockEngine := NewMockEngine(ct)
	svc := New(mockCar, mockEngine)

	mockCar.EXPECT().CarGetByBrand(ctx, "bmw").Return([]*models.Car{}, errors.New("error from datastore layer"))

	res, err := svc.GetByBrand(ctx, "bmw", true)
	assert.Equal(t, errors.New("error from datastore layer"), err)
	assert.Equal(t, []*models.Car{}, res)
}

// TestService_GetByBrandEngineError method is used for testing GetByBrand method in Service layer.
func TestService_GetByBrandEngineError(t *testing.T) {
	ctx := context.Background()
	ct := gomock.NewController(t)
	mockCar := NewMockCar(ct)
	mockEngine := NewMockEngine(ct)
	svc := New(mockCar, mockEngine)

	id := uuid.New()
	car := []*models.Car{{ID: id, Name: "b3", Year: 2000, Brand: "bmw", Fuel: "petrol", Engine: models.Engine{ID: id, Displacement: 200,
		NoOfCylinders: 3, Range: 0}}}

	mockCar.EXPECT().CarGetByBrand(ctx, "bmw").Return(car, nil)
	mockEngine.EXPECT().EngineGetByID(ctx, id).Return(&models.Engine{}, errors.New("error from datastore layer"))

	res, err := svc.GetByBrand(ctx, "bmw", true)
	assert.Equal(t, errors.New("error from datastore layer"), err)
	assert.Equal(t, []*models.Car{}, res)
}

// TestService_GetById method is used for testing GetById method in Service layer.
func TestService_GetByID(t *testing.T) {
	ctx := context.Background()
	car := models.Car{ID: uuid.New(), Name: "Model X", Year: 2018, Brand: "Ferrari", Fuel: "Electric", Engine: models.Engine{Range: 100}}
	id := uuid.New()

	ct := gomock.NewController(t)
	mockCar := NewMockCar(ct)
	mockEngine := NewMockEngine(ct)
	svc := New(mockCar, mockEngine)

	mockCar.EXPECT().CarGetByID(context.TODO(), id).Return(&car, nil)
	mockEngine.EXPECT().EngineGetByID(ctx, car.Engine.ID).Return(&car.Engine, nil)
	_, err := svc.GetByID(ctx, id)

	assert.Equal(t, nil, err)
}

// TestService_GetByIdEngineError method is used for testing GetById method in Service layer.
func TestService_GetByIDEngineError(t *testing.T) {
	ctx := context.Background()
	car := models.Car{ID: uuid.New(), Name: "Model X", Year: 2018, Brand: "Ferrari", Fuel: "Electric", Engine: models.Engine{Range: 100}}
	id := uuid.New()

	ct := gomock.NewController(t)
	mockCar := NewMockCar(ct)
	mockEngine := NewMockEngine(ct)
	svc := New(mockCar, mockEngine)

	mockCar.EXPECT().CarGetByID(context.TODO(), id).Return(&car, nil)
	mockEngine.EXPECT().EngineGetByID(ctx, car.Engine.ID).Return(&models.Engine{}, errors.New("error from engine in datastore layer"))
	_, err := svc.GetByID(ctx, id)

	assert.Equal(t, errors.New("error from engine in datastore layer"), err)
}

// TestService_GetByIdEngineError method is used for testing GetById method in Service layer.
func TestService_GetByIDCarError(t *testing.T) {
	ctx := context.Background()
	id := uuid.New()

	ct := gomock.NewController(t)
	mockCar := NewMockCar(ct)
	mockEngine := NewMockEngine(ct)
	svc := New(mockCar, mockEngine)

	mockCar.EXPECT().CarGetByID(context.TODO(), id).Return(&models.Car{}, errors.New("error from engine in datastore layer"))
	_, err := svc.GetByID(ctx, id)

	assert.Equal(t, errors.New("error from engine in datastore layer"), err)
}

// TestService_Update method is used for testing Update method in Service layer.
func TestService_Update(t *testing.T) {
	ctx := context.Background()
	ct := gomock.NewController(t)
	mockCar := NewMockCar(ct)
	mockEngine := NewMockEngine(ct)
	svc := New(mockCar, mockEngine)

	id := uuid.New()
	car := models.Car{ID: id, Name: "b3", Year: 2001, Brand: "bmw", Fuel: "petrol", Engine: models.Engine{ID: id, Displacement: 500,
		NoOfCylinders: 200, Range: 0}}

	mockCar.EXPECT().CarUpdate(ctx, &car).Return(&car, nil)
	mockCar.EXPECT().CarGetByID(ctx, id).Return(&car, nil)
	mockEngine.EXPECT().EngineUpdate(ctx, &car.Engine).Return(&car.Engine, nil)

	res, err := svc.Update(ctx, &car)

	assert.Equal(t, &car, res)
	assert.Equal(t, nil, err)
}

// TestService_UpdateInvalidCar method is used for testing Update method in Service layer.
func TestService_UpdateInvalidBrand(t *testing.T) {
	ctx := context.Background()
	ct := gomock.NewController(t)
	mockCar := NewMockCar(ct)
	mockEngine := NewMockEngine(ct)
	svc := New(mockCar, mockEngine)

	id := uuid.New()
	car := models.Car{ID: id, Name: "b3", Year: 2001, Brand: "maruthi", Fuel: "petrol", Engine: models.Engine{ID: id, Displacement: 500,
		NoOfCylinders: 200, Range: 0}}

	res, err := svc.Update(ctx, &car)

	assert.Equal(t, &models.Car{}, res)
	assert.Equal(t, errors.New("invalid brand"), err)
}

// TestService_UpdateCarError method is used for testing Update method in Service layer.
func TestService_UpdateCarError(t *testing.T) {
	ctx := context.Background()
	ct := gomock.NewController(t)

	id := uuid.New()
	mockCar := NewMockCar(ct)
	mockEngine := NewMockEngine(ct)
	svc := New(mockCar, mockEngine)

	car := models.Car{ID: id, Name: "b3", Year: 2001, Brand: "bmw", Fuel: "petrol", Engine: models.Engine{ID: id, Displacement: 500,
		NoOfCylinders: 200, Range: 0}}

	mockCar.EXPECT().CarUpdate(ctx, &car).Return(&models.Car{}, errors.New("error from car in datastore layer"))

	res, err := svc.Update(ctx, &car)

	assert.Equal(t, &models.Car{}, res)
	assert.Equal(t, errors.New("error from car in datastore layer"), err)
}

// TestService_UpdateIdError method is used for testing Update method in Service layer.
func TestService_UpdateIdError(t *testing.T) {
	ctx := context.Background()
	ct := gomock.NewController(t)
	mockCar := NewMockCar(ct)
	mockEngine := NewMockEngine(ct)
	svc := New(mockCar, mockEngine)

	id := uuid.New()
	car := models.Car{ID: id, Name: "b3", Year: 2001, Brand: "bmw", Fuel: "petrol", Engine: models.Engine{ID: id, Displacement: 500,
		NoOfCylinders: 200, Range: 0}}

	mockCar.EXPECT().CarUpdate(ctx, &car).Return(&car, nil)
	mockCar.EXPECT().CarGetByID(ctx, id).Return(&models.Car{}, errors.New("error from car in datastore layer"))

	res, err := svc.Update(ctx, &car)

	assert.Equal(t, &models.Car{}, res)
	assert.Equal(t, errors.New("error from car in datastore layer"), err)
}

// TestService_UpdateEngineError method is used for testing Update method in Service layer.
func TestService_UpdateEngineError(t *testing.T) {
	ctx := context.Background()
	ct := gomock.NewController(t)
	mockCar := NewMockCar(ct)
	mockEngine := NewMockEngine(ct)
	svc := New(mockCar, mockEngine)

	id := uuid.New()
	car := models.Car{ID: id, Name: "b3", Year: 2001, Brand: "bmw", Fuel: "petrol", Engine: models.Engine{ID: id, Displacement: 500,
		NoOfCylinders: 200, Range: 0}}

	mockCar.EXPECT().CarUpdate(ctx, &car).Return(&car, nil)
	mockCar.EXPECT().CarGetByID(ctx, id).Return(&car, nil)
	mockEngine.EXPECT().EngineUpdate(ctx, &car.Engine).Return(&models.Engine{}, errors.New("error from datastore layer"))

	res, err := svc.Update(ctx, &car)

	assert.Equal(t, &models.Car{}, res)
	assert.Equal(t, errors.New("error from datastore layer"), err)
}

// TestService_Delete method is used for testing Delete method in Service layer.
func TestService_Delete(t *testing.T) {
	ctx := context.Background()
	car := models.Car{ID: uuid.New(), Name: "Model X", Year: 2018, Brand: "Ferrari", Fuel: "Electric", Engine: models.Engine{Range: 100}}
	id := uuid.New()

	ct := gomock.NewController(t)
	mockCar := NewMockCar(ct)
	mockEngine := NewMockEngine(ct)
	svc := New(mockCar, mockEngine)

	mockCar.EXPECT().CarGetByID(context.TODO(), id).Return(&car, nil)
	mockEngine.EXPECT().EngineDelete(ctx, car.Engine.ID).Return(nil)
	mockCar.EXPECT().CarDelete(ctx, id).Return(nil)
	err := svc.Delete(ctx, id)

	assert.Equal(t, nil, err)
}

// TestService_DeleteIdError method is used for testing Delete method in Service layer.
func TestService_DeleteIdError(t *testing.T) {
	ctx := context.Background()
	id := uuid.New()

	ct := gomock.NewController(t)
	mockCar := NewMockCar(ct)
	mockEngine := NewMockEngine(ct)
	svc := New(mockCar, mockEngine)

	mockCar.EXPECT().CarGetByID(context.TODO(), id).Return(&models.Car{}, errors.New("error from getById in datastore layer"))
	err := svc.Delete(ctx, id)

	assert.Equal(t, errors.New("error from getById in datastore layer"), err)
}

// TestService_DeleteCarError method is used for testing Delete method in Service layer.
func TestService_DeleteCarError(t *testing.T) {
	ctx := context.Background()
	car := models.Car{ID: uuid.New(), Name: "Model X", Year: 2018, Brand: "Ferrari", Fuel: "Electric", Engine: models.Engine{Range: 100}}
	id := uuid.New()

	ct := gomock.NewController(t)
	mockCar := NewMockCar(ct)
	mockEngine := NewMockEngine(ct)
	svc := New(mockCar, mockEngine)

	mockCar.EXPECT().CarGetByID(context.TODO(), id).Return(&car, nil)
	mockEngine.EXPECT().EngineDelete(ctx, car.Engine.ID).Return(errors.New("error from car in datastore layer"))

	err := svc.Delete(ctx, id)

	assert.Equal(t, errors.New("error from car in datastore layer"), err)
}

// TestService_DeleteEngineError method is used for testing Delete method in Service layer.
func TestService_DeleteEngineError(t *testing.T) {
	ctx := context.Background()
	ct := gomock.NewController(t)
	mockCar := NewMockCar(ct)
	mockEngine := NewMockEngine(ct)
	svc := New(mockCar, mockEngine)

	car := models.Car{ID: uuid.New(), Name: "Model X", Year: 2018, Brand: "Ferrari", Fuel: "Electric", Engine: models.Engine{Range: 100}}
	id := uuid.New()

	mockCar.EXPECT().CarGetByID(context.TODO(), id).Return(&car, nil)
	mockEngine.EXPECT().EngineDelete(ctx, car.Engine.ID).Return(nil)
	mockCar.EXPECT().CarDelete(ctx, id).Return(errors.New("error from engine in datastore layer"))
	err := svc.Delete(ctx, id)

	assert.Equal(t, errors.New("error from engine in datastore layer"), err)
}

// TestCheck method is used for testing check function in Service layer.
func TestValidation(t *testing.T) {
	testcases := []struct {
		desc  string
		input models.Car
		err   error
	}{
		{"invalid year",
			models.Car{
				ID:    uuid.MustParse("8f443772-132b-4ae5-9f8f-9960649b3fb4"),
				Name:  "b3",
				Year:  2025,
				Brand: "bmw",
				Fuel:  "petrol",
				Engine: models.Engine{
					ID:            uuid.MustParse("8f443772-132b-4ae5-9f8f-9960649b3fb4"),
					Displacement:  200,
					Range:         0,
					NoOfCylinders: 4,
				},
			},
			errors.New("invalid Year"),
		},
		{"petrol car doesn't have range",
			models.Car{
				ID:    uuid.MustParse("8f443772-132b-4ae5-9f8f-9960649b3fb4"),
				Name:  "b3",
				Year:  2005,
				Brand: "tesla",
				Fuel:  "petrol",
				Engine: models.Engine{
					ID:            uuid.MustParse("8f443772-132b-4ae5-9f8f-9960649b3fb4"),
					Displacement:  200,
					Range:         90,
					NoOfCylinders: 4,
				},
			}, errors.New("petrol car doesn't have range"),
		},
		{"diesel car doesn't have range",
			models.Car{
				ID:    uuid.MustParse("8f443772-132b-4ae5-9f8f-9960649b3fb4"),
				Name:  "b3",
				Year:  2005,
				Brand: "tesla",
				Fuel:  "diesel",
				Engine: models.Engine{
					ID:            uuid.MustParse("8f443772-132b-4ae5-9f8f-9960649b3fb4"),
					Displacement:  200,
					Range:         90,
					NoOfCylinders: 4,
				},
			}, errors.New("diesel car doesn't have range"),
		},
		{"invalid brand",
			models.Car{
				ID:    uuid.MustParse("8f443772-132b-4ae5-9f8f-9960649b3fb4"),
				Name:  "b3",
				Year:  2005,
				Brand: "Honda",
				Fuel:  "diesel",
				Engine: models.Engine{
					ID:            uuid.MustParse("8f443772-132b-4ae5-9f8f-9960649b3fb4"),
					Displacement:  200,
					Range:         0,
					NoOfCylinders: 4,
				},
			}, errors.New("invalid brand")},
		{"invalid fuel",
			models.Car{
				ID:    uuid.MustParse("8f443772-132b-4ae5-9f8f-9960649b3fb4"),
				Name:  "b3",
				Year:  2005,
				Brand: "tesla",
				Fuel:  "battery",
				Engine: models.Engine{
					ID:            uuid.MustParse("8f443772-132b-4ae5-9f8f-9960649b3fb4"),
					Displacement:  200,
					Range:         50,
					NoOfCylinders: 4,
				},
			}, errors.New("invalid fuel")},
		{"electric car",
			models.Car{
				ID:    uuid.MustParse("8f443772-132b-4ae5-9f8f-9960649b3fb4"),
				Name:  "b3",
				Year:  2005,
				Brand: "tesla",
				Fuel:  "electric",
				Engine: models.Engine{
					ID:            uuid.MustParse("8f443772-132b-4ae5-9f8f-9960649b3fb4"),
					Displacement:  200,
					Range:         0,
					NoOfCylinders: 4,
				},
			}, errors.New("electric car doesn't have displacement and no.of.cylinders")},
	}
	for _, tc := range testcases {
		err := validation(tc.input)
		assert.Equal(t, err, tc.err)
	}
}
