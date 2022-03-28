package crud

import (
	"bytes"
	"database/sql"
	"encoding/json"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"io"
	"log"
	"model"
	"net/http"
	"net/http/httptest"
	"testing"
)

func DbConn() (db *sql.DB, err error) {
	dbDriver := "mysql"
	dbUser := "manasa"
	dbPass := "Manasa@2210"
	dbname := "Dealership"
	db, err = sql.Open(dbDriver, dbUser+":"+dbPass+"@tcp(localhost:3306)"+"/"+dbname)
	if err != nil {
		return nil, err
	}
	return db, nil
}

// Test_Create is used to test the Create function
func Test_Create(t *testing.T) {
	db, err := DbConn()
	if err != nil {
		log.Printf("unexpected error %v", err)
		return
	}
	d := New(db)
	testcases := []struct {
		desc   string
		input  model.Car
		output string
	}{
		{desc: "success case", input: model.Car{Id: uuid.New(), Name: "AB", Year: 2000, Brand: model.Tesla, Fuel: model.Electric,
			Engine: model.Engine{Range: 4}}, output: "data successfully entered"},
		//{desc: "year less than 1980", input: model.Car{Id: uuid.New(), Name: "A", Year: 1970, Brand: model.Porsche, Fuel: model.Petrol,
		//	Engine: model.Engine{Displacement: 56.7, NoOfCylinders: 2, Range: 4}}, output: "year must be between 1980 and 2022"},
		//{desc: "year greater than 2022", input: model.Car{Id: uuid.New(), Name: "ABC", Year: 2040, Brand: model.Porsche, Fuel: model.Petrol,
		//Engine: model.Engine{Displacement: 56.7, NoOfCylinders: 2, Range: 4}}, output: "year must be between 1980 and 2022"},
		//{desc: "invalid brand ", input: model.Car{Id: uuid.New(), Name: "AB", Year: 2000, Brand: "Nano", Fuel: model.Diesel,
		//Engine: model.Engine{Displacement: 56.7, NoOfCylinders: 2, Range: 4}}, output: "Status Bad Request"},
		//{desc: "invalid fuel ", input: model.Car{Id: uuid.New(), Name: "AB", Year: 2000, Brand: model.Porsche, Fuel: "Water",
		//Engine: model.Engine{Displacement: 56.7, NoOfCylinders: 2, Range: 4}}, output: "Status Bad Request"},
	}

	for _, tc := range testcases {
		data, err := json.Marshal(tc.input)
		if err != nil {
			t.Errorf("unexpected error %v", err)
		}
		req := httptest.NewRequest(http.MethodPost, "/cars", bytes.NewReader(data))
		w := httptest.NewRecorder()

		d.Create(w, req)
		resp := w.Result()
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			t.Errorf("unexpected error %v", err)
		}
		if !(assert.Equal(t, string(body), tc.output)) {
			t.Errorf("failed %v expected %v but got %v", tc.desc, tc.output, string(body))
		}
	}
}

// Test_GetByBrand is used to test the GetByBrand function
func Test_GetByBrand(t *testing.T) {
	db, err := DbConn()
	if err != nil {
		log.Printf("unexpected error %v", err)
		return
	}
	d := New(db)
	testcases := []struct {
		desc   string
		input  string
		output []model.Car
	}{
		{desc: "success case with brand and engine", input: "/car?brand=Tesla&engine=included", output: []model.Car{{Id: uuid.MustParse("8f14a65f-3032-42c8-a196-1cf66d11b932"),
			Name: "AB", Year: 2000, Brand: model.Tesla, Fuel: model.Electric, Engine: model.Engine{Displacement: 56.7, NoOfCylinders: 2, Range: 4}}}},
		{"invalid input", "/car?brand=", nil},
		{"success case with brand", "/car?brand:Tesla", []model.Car{{uuid.MustParse("8f14a65f-3032-42c8-a196-1cf66d11b932"),
			"AB", 2000, model.Tesla, model.Petrol, model.Engine{}}}},
	}
	for _, tc := range testcases {
		req := httptest.NewRequest(http.MethodGet, tc.input, nil)
		w := httptest.NewRecorder()
		d.GetByBrand(w, req)
		resp := w.Result()

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			t.Errorf("unexpected error %v", err)
		}
		if assert.Equal(t, string(body), tc.output) {
			t.Errorf("expected %v but got %v", tc.output, string(body))
		}
	}
}

// Test_GetById is used to test the GetById function
func Test_GetById(t *testing.T) {
	db, err := DbConn()
	if err != nil {
		log.Printf("unexpected error %v", err)
		return
	}
	d := New(db)
	testcases := []struct {
		desc   string
		input  string
		output string
	}{
		{"success case", "/car/a24a6ea4-ce75-4665-a070-57453082c256",
			"row with given ID is fetched"},
		{"invalid id", "/car/id/1", "Status bad request"},
		{"invalid id", "/car/id", "Status bad request"},
	}
	for _, v := range testcases {
		req := httptest.NewRequest(http.MethodGet, v.input, nil)
		w := httptest.NewRecorder()
		d.GetById(w, req)
		resp := w.Result()
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			t.Errorf("unexpected error %v", err)
		}
		if assert.Equal(t, string(body), v.output) {
			t.Errorf("expected %v but got %v", v.output, string(body))
		}
	}
}

// Test_Update is used to test the Update function
func Test_Update(t *testing.T) {
	db, err := DbConn()
	if err != nil {
		log.Printf("unexpected error %v", err)
		return
	}
	d := New(db)
	testcases := []struct {
		desc   string
		input  model.Car
		output string
	}{
		{desc: "success case", input: model.Car{Id: uuid.New(), Name: "AB", Year: 2000, Brand: model.Tesla, Fuel: model.Electric,
			Engine: model.Engine{Range: 4}}, output: "data successfully entered"},
		{desc: "year less than 1980", input: model.Car{Id: uuid.New(), Name: "A", Year: 1970, Brand: model.Porsche, Fuel: model.Petrol,
			Engine: model.Engine{Displacement: 56.7, NoOfCylinders: 2, Range: 4}}, output: "year must be between 1980 and 2022"},
		{desc: "year greater than 2022", input: model.Car{Id: uuid.New(), Name: "ABC", Year: 2040, Brand: model.Porsche, Fuel: model.Petrol,
			Engine: model.Engine{Displacement: 56.7, NoOfCylinders: 2, Range: 4}}, output: "year must be between 1980 and 2022"},
		{desc: "invalid brand ", input: model.Car{Id: uuid.New(), Name: "AB", Year: 2000, Brand: "Nano", Fuel: model.Diesel,
			Engine: model.Engine{Displacement: 56.7, NoOfCylinders: 2, Range: 4}}, output: "enter a valid brand"},
		{desc: "invalid fuel ", input: model.Car{Id: uuid.New(), Name: "AB", Year: 2000, Brand: model.Porsche, Fuel: "Water",
			Engine: model.Engine{Displacement: 56.7, NoOfCylinders: 2, Range: 4}}, output: "enter a valid fuel"},
	}
	for _, tc := range testcases {
		data, err := json.Marshal(tc.input)
		if err != nil {
			t.Errorf("unexpected error %v", err)
		}
		req := httptest.NewRequest(http.MethodPut, "/car", bytes.NewReader(data))
		w := httptest.NewRecorder()
		d.Update(w, req)
		resp := w.Result()
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			t.Errorf("unexpected error %v", err)
		}
		if assert.Equal(t, string(body), tc.output) {
			t.Errorf("expected %v but got %v", tc.output, string(body))
		}
	}
}

// Test_Delete is used to test the Delete function
func Test_Delete(t *testing.T) {
	db, err := DbConn()
	if err != nil {
		log.Printf("unexpected error %v", err)
		return
	}
	d := New(db)
	testcases := []struct {
		desc   string
		input  string
		output string
	}{
		{"success case", "/car/id/a24a6ea4-ce75-4665-a070-57453082c256", "row with given ID is deleted"},
		{"success case", "/car/id/1", "invalid request"},
		{"invalid input parameters", "/car/", "Status bad request"},
	}
	for _, tc := range testcases {
		req := httptest.NewRequest(http.MethodDelete, tc.input, nil)
		w := httptest.NewRecorder()
		d.Delete(w, req)
		resp := w.Result()
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			t.Errorf("unexpected error %v", err)
		}
		if assert.Equal(t, string(body), tc.output) {
			t.Errorf("expected %v but got %v", tc.output, string(body))
		}
	}
}
