package handlers

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strconv"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"

	"manasa/models"
)

// TestHandler_Delete is used to test Delete method in handler layer.
func TestHandler_Delete(t *testing.T) {
	testcases := []struct {
		desc     string
		id       uuid.UUID
		err      error
		expected int
	}{
		{"success case", uuid.MustParse("8f443772-132b-4ae5-9f8f-9960649b3fb4"),
			nil, http.StatusNoContent},
		{"invalid id", uuid.MustParse("8f443772-132b-4ae5-9f8f-9960649b3fb6"),
			errors.New("error from service layer"), http.StatusInternalServerError},
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	m := NewMockService(ctrl)

	for i, tc := range testcases {
		req := httptest.NewRequest(http.MethodGet, "https://car", nil)
		r := mux.SetURLVars(req, map[string]string{"id": tc.id.String()})
		w := httptest.NewRecorder()

		m.EXPECT().Delete(gomock.Any(), tc.id).Return(tc.err)
		h := New(m)
		h.Delete(w, r)

		if w.Code != tc.expected {
			t.Errorf("[Test %d]Failed. Got %v Expected %v/n", i, w.Code, tc.expected)
		}
	}
}

// TestMockService_DeleteError is used for testing error cases in create.
func TestHandler_DeleteError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := NewMockService(ctrl)

	testcases := []struct {
		desc   string
		id     string
		status int
	}{
		{"wrong id", "123", http.StatusBadRequest},
	}

	for _, tc := range testcases {
		req := httptest.NewRequest(http.MethodDelete, "https://car/id", nil)
		w := httptest.NewRecorder()

		h := New(m)
		h.Delete(w, req)

		assert.Equal(t, tc.status, w.Code)
	}
}

// TestHandler_GetByID is used to test GetById method in handler layer.
func TestHandler_GetByID(t *testing.T) {
	id := uuid.New()
	output := models.Car{Name: "S-200", Year: 2018, Brand: models.Ferrari, Fuel: models.Petrol,
		Engine: models.Engine{Displacement: 2500, NoOfCylinders: 6}}

	testcases := []struct {
		desc       string
		statusCode int
		output     models.Car
		err        error
	}{
		{"Success", http.StatusOK, output, nil},
		{"invalid id", http.StatusInternalServerError, models.Car{}, errors.New("error from service layer")},
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	m := NewMockService(ctrl)

	for _, tc := range testcases {
		req := httptest.NewRequest(http.MethodGet, "https://car/id", nil)
		r := mux.SetURLVars(req, map[string]string{"id": id.String()})
		w := httptest.NewRecorder()

		m.EXPECT().GetByID(gomock.Any(), id).Return(&tc.output, tc.err)
		h := New(m)
		h.GetByID(w, r)

		assert.Equal(t, w.Code, tc.statusCode)
	}
}

// TestMockService_GetByIDError is used for testing error cases in create.
func TestHandler_GetByIDError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := NewMockService(ctrl)

	testcases := []struct {
		desc   string
		id     string
		status int
	}{
		{"wrong id", "123", http.StatusBadRequest},
	}

	for _, tc := range testcases {
		req := httptest.NewRequest(http.MethodGet, "https://car/id", nil)
		w := httptest.NewRecorder()

		h := New(m)
		h.GetByID(w, req)

		assert.Equal(t, tc.status, w.Code)
	}
}

// TestHandler_GetByBrand is used to test GetByBrand method in handler layer.
func TestHandler_GetByBrand(t *testing.T) {
	var car1 = models.Car{ID: uuid.MustParse("8f443772-132b-4ae5-9f8f-9960649b3fb4"), Name: "b3", Year: 2003,
		Brand: "BMW", Fuel: "petrol", Engine: models.Engine{Displacement: 200, Range: 0, NoOfCylinders: 3,
			ID: uuid.MustParse("8f443772-132b-4ae5-9f8f-9960649b3fb4")}}

	ctx := context.Background()
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	m := NewMockService(ctrl)

	testcases := []struct {
		desc           string
		brand          string
		engine         bool
		resBody        []*models.Car
		err            error
		expectedStatus int
	}{
		{"success case", "BMW", true, []*models.Car{&car1}, nil, http.StatusOK},
		{"invalid case", "BMW", true, []*models.Car{}, errors.New("error from service layer"), http.StatusInternalServerError},
	}

	for i, tc := range testcases {
		req := httptest.NewRequest(http.MethodGet, "https://car?brand="+tc.brand+"&isEngine="+strconv.FormatBool(tc.engine), nil)
		w := httptest.NewRecorder()

		m.EXPECT().GetByBrand(ctx, tc.brand, tc.engine).Return(tc.resBody, tc.err)
		h := New(m)
		h.GetByBrand(w, req)

		if !reflect.DeepEqual(w.Code, tc.expectedStatus) {
			t.Errorf("[Test %d]Failed Got %v Expected %v", i, w.Code, tc.expectedStatus)
		}
	}
}

// TestMockService_GetByBrandError is used for testing error cases in create.
func TestHandler_GetByBrandError(t *testing.T) {
	ctrl := gomock.NewController(t)

	defer ctrl.Finish()

	m := NewMockService(ctrl)

	testcases := []struct {
		desc           string
		brand          string
		engine         string
		expectedStatus int
	}{
		{"invalid isEngine", "BMW", "tr", http.StatusBadRequest},
	}

	for _, tc := range testcases {
		req := httptest.NewRequest(http.MethodGet, "https://car?brand="+tc.brand+"&isEngine="+tc.engine, nil)
		w := httptest.NewRecorder()

		h := New(m)
		h.GetByBrand(w, req)

		assert.Equal(t, tc.expectedStatus, w.Code)
	}
}

// TestHandler_Update is used to test Update method in handler layer.
func TestHandler_Update(t *testing.T) {
	var car1 = models.Car{ID: uuid.MustParse("8f443772-132b-4ae5-9f8f-9960649b3fb4"), Name: "b3", Year: 2003,
		Brand: "BMW", Fuel: "petrol", Engine: models.Engine{Displacement: 200, Range: 0, NoOfCylinders: 3,
			ID: uuid.MustParse("8f443772-132b-4ae5-9f8f-9960649b3fb4")}}

	var car2 = models.Car{ID: uuid.MustParse("8f443772-132b-4ae5-9f8f-9960649b3fb7"), Name: "b3", Year: 2003,
		Brand: "BMW", Fuel: "petrol", Engine: models.Engine{Displacement: 200, Range: 0, NoOfCylinders: 3,
			ID: uuid.MustParse("8f443772-132b-4ae5-9f8f-9960649b3fb7")}}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := NewMockService(ctrl)

	testcases := []struct {
		desc               string
		id                 uuid.UUID
		reqBody            models.Car
		resBody            models.Car
		err                error
		expectedStatusCode int
	}{
		{"success case", car1.ID, car1, car1, nil, http.StatusOK},
		{"invalid id", car2.ID, car2, models.Car{},
			errors.New("error from service layer"), http.StatusInternalServerError},
	}

	for i, tc := range testcases {
		body, err := json.Marshal(tc.reqBody)
		if err != nil {
			fmt.Println(err)
		}

		req := httptest.NewRequest(http.MethodPut, "https://car/id", bytes.NewBuffer(body))
		r := mux.SetURLVars(req, map[string]string{"id": tc.id.String()})
		w := httptest.NewRecorder()

		m.EXPECT().Update(gomock.Any(), &tc.reqBody).Return(&tc.resBody, tc.err)
		h := New(m)
		h.Update(w, r)

		if w.Code != tc.expectedStatusCode {
			t.Errorf("[TEST%d]Failed. Got %v\tExpected %v\n", i, w.Code, tc.expectedStatusCode)
		}
	}
}

// TestMockService_UpdateError is used for testing error cases in create.
func TestMockService_UpdateError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := NewMockService(ctrl)

	testcases := []struct {
		desc   string
		id     string
		input  []byte
		status int
	}{
		{"missing parameters", (uuid.New()).String(), []byte(`[{"name": "bmw 2", "year": 2000}]`), http.StatusBadRequest},
		{"invalid id", "hello", []byte{}, http.StatusBadRequest},
	}

	for _, tc := range testcases {
		req := httptest.NewRequest(http.MethodPost, "https://car/"+tc.id, bytes.NewBuffer(tc.input))
		w := httptest.NewRecorder()

		h := New(m)
		h.Update(w, req)

		assert.Equal(t, tc.status, w.Code)
	}
}

// TestHandler_Create is used to test Create method in handler layer.
func TestHandler_Create(t *testing.T) {
	var car1 = models.Car{ID: uuid.MustParse("8f443772-132b-4ae5-9f8f-9960649b3fb4"), Name: "b3", Year: 2003,
		Brand: "BMW", Fuel: "petrol", Engine: models.Engine{Displacement: 200, Range: 0, NoOfCylinders: 3,
			ID: uuid.MustParse("8f443772-132b-4ae5-9f8f-9960649b3fb4")}}

	var car2 = models.Car{ID: uuid.MustParse("8f443772-132b-4ae5-9f8f-9960649b3fb7"), Name: "b3", Year: 2003,
		Brand: "BMW", Fuel: "petrol", Engine: models.Engine{Displacement: 200, Range: 0, NoOfCylinders: 3,
			ID: uuid.MustParse("8f443772-132b-4ae5-9f8f-9960649b3fb7")}}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := NewMockService(ctrl)

	testcases := []struct {
		desc               string
		reqBody            models.Car
		resBody            models.Car
		err                error
		expectedStatusCode int
	}{
		{"success case", car1, car1, nil, http.StatusCreated},
		{"invalid id", car2, models.Car{}, errors.New("error from service layer"), http.StatusInternalServerError},
	}

	for i, tc := range testcases {
		body, err := json.Marshal(tc.reqBody)
		if err != nil {
			fmt.Println(err)
		}

		req := httptest.NewRequest(http.MethodPost, "https://car", bytes.NewBuffer(body))
		w := httptest.NewRecorder()

		m.EXPECT().Create(gomock.Any(), &tc.reqBody).Return(&tc.resBody, tc.err)
		h := New(m)
		h.Create(w, req)

		resp := w.Result()

		_, err = io.ReadAll(resp.Body)
		if err != nil {
			t.Errorf("error in reading output %v", err)
		}

		resp.Body.Close()

		if w.Code != tc.expectedStatusCode {
			t.Errorf("[TEST%d]Failed. Got %v\tExpected %v\n", i, w.Code, tc.expectedStatusCode)
		}
	}
}

// TestMockService_CreateError is used for testing error cases in create.
func TestMockService_CreateError(t *testing.T) {
	testcases := []struct {
		desc   string
		input  []byte
		status int
	}{
		{"missing parameters", []byte(`[{"name": "bmw 2", "year": 2000}]`), http.StatusBadRequest},
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := NewMockService(ctrl)

	for _, tc := range testcases {
		req := httptest.NewRequest(http.MethodPost, "https://car", bytes.NewBuffer(tc.input))
		w := httptest.NewRecorder()

		h := New(m)
		h.Create(w, req)

		assert.Equal(t, tc.status, w.Code)
	}
}
