package crud

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"io/ioutil"
	_ "io/ioutil"
	"log"
	"model"
	"net/http"
)

type dB struct {
	DB *sql.DB
}

func New(db *sql.DB) dB {
	return dB{DB: db}
}

// Create is used for Inserting values to database
func (d dB) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	res, _ := ioutil.ReadAll(r.Body)
	data := model.Car{}

	err := json.Unmarshal(res, &data)
	if err != nil {
		_ = fmt.Errorf("unexpected error %v", err)
	}
	if data.Year < 1980 || data.Year > 2022 {
		_, err := w.Write([]byte("year must be between 1980 and 2022"))
		if err != nil {
			_ = fmt.Errorf("unexpected error %v", err)
		}
	}

	if data.Brand != "Tesla" && data.Brand != "Porsche" && data.Brand != "Ferrari" && data.Brand != "Mercedes" && data.Brand != "BMW" {
		w.WriteHeader(http.StatusBadRequest)
		_, err := w.Write([]byte("Status Bad Request"))
		if err != nil {
			_ = fmt.Errorf("unexpected error %v", err)
		}
	}
	if data.Fuel != "Petrol" && data.Fuel != "Diesel" && data.Fuel != "Electric" {
		w.WriteHeader(http.StatusBadRequest)
		_, err := w.Write([]byte("Status Bad Request"))
		if err != nil {
			_ = fmt.Errorf("unexpected error %v", err)
		}
	}
	if data.Fuel == "Electric" {
		data.Engine.Displacement = 0
		data.Engine.NoOfCylinders = 0
		if data.Engine.Range < 0 {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}
	if data.Fuel == "Petrol" || data.Fuel == "Diesel" {
		if data.Engine.Displacement <= 0 {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if data.Engine.NoOfCylinders <= 0 {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		data.Engine.Range = 0
	}

	data.Id = uuid.New()

	_, err = d.DB.Exec("INSERT INTO car (Id,Name,Year,Brand,FuelType)Values(?,?,?,?,?)", data.Id, data.Name, data.Year, data.Brand, data.Fuel)
	if err != nil {
		_ = fmt.Errorf("unexpected error %v", err)
	}
	if err != nil {
		// todo use errorf to log errors
		_ = fmt.Errorf("unexpected error %v", err)
	}
	_, err = d.DB.Exec("INSERT INTO engine(Car_id,Displacement,No_of_cylinders,`Range`)VALUES(?,?,?,?)", data.Id, data.Engine.Displacement, data.Engine.NoOfCylinders, data.Engine.Range)
	if err != nil {
		_ = fmt.Errorf("unexpected error %v", err)
	}

	// todo handler errors
	err = json.NewEncoder(w).Encode(data)
	if err != nil {
		_ = fmt.Errorf("unexpected error %v", err)
	}

}

// GetByBrand is used for getting values from database
func (d dB) GetByBrand(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	brand := r.URL.Query().Get("brand")
	engine := r.URL.Query().Get("engine")
	brand = string(bytes.TrimSpace([]byte(brand)))
	engine = string(bytes.TrimSpace([]byte(engine)))
	if brand != "" && engine == "" {
		_, err := w.Write([]byte(" data with Brand " + brand))
		if err != nil {
			_ = fmt.Errorf("unexpected error %v", err)
		}
		queryCar, err := d.DB.Query("Select * from car where Brand=?", brand)
		if err != nil {
			_ = fmt.Errorf("unexpected error %v", err)
		}
		var cars model.Car
		var res []model.Car

		for queryCar.Next() {
			err := queryCar.Scan(&cars.Id, &cars.Name, &cars.Year, &cars.Brand, &cars.Fuel)
			if err != nil {
				_ = fmt.Errorf("unexpected error %v", err)
			}
			res = append(res, cars)
			fmt.Println(res)
		}
		output, err := json.Marshal(res)
		if err != nil {
			_ = fmt.Errorf("unexpected error %v", err)
		}
		_, err = w.Write(output)
		if err != nil {
			_ = fmt.Errorf("unexpected error %v", err)
		}
	} else if len(brand) != 0 && len(engine) != 0 {
		queryCar, err := d.DB.Query("Select * from car where Brand=?", brand)
		if err != nil {
			_ = fmt.Errorf("unexpected error %v", err)
		}
		var cars model.Car
		var res []model.Car

		for queryCar.Next() {
			err := queryCar.Scan(&cars.Id, &cars.Name, &cars.Year, &cars.Brand, &cars.Fuel)
			if err != nil {
				_ = fmt.Errorf("unexpected error %v", err)
			}

			d.DB.QueryRow("Select Displacement,No_of_cylinders,`Range` from engine where Car_id=?", cars.Id).
				Scan(&cars.Engine.Displacement, &cars.Engine.NoOfCylinders, &cars.Engine.Range)
			res = append(res, cars)

		}
		output, err := json.Marshal(res)
		if err != nil {
			_ = fmt.Errorf("unexpected error %v", err)
		}
		_, err = w.Write(output)
		if err != nil {
			_ = fmt.Errorf("unexpected error %v", err)
		}
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}

// todo fix linters

// GetById is used for getting values from database using Id
func (d dB) GetById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	param := mux.Vars(r)
	id := param["Id"]
	res := model.Car{}
	queryCar := d.DB.QueryRow("Select * from car where Id=?", id)
	err := queryCar.Scan(&res.Id, &res.Name, &res.Year, &res.Brand, &res.Fuel)
	if err != nil {
		_, err = w.Write([]byte("invalid parameters"))
		if err != nil {
			_ = fmt.Errorf("unexpected error %v", err)
		}
		return
	}

	d.DB.QueryRow("Select Displacement,No_of_cylinders,`Range` from engine where Car_id=?", id).Scan(&res.Engine.Displacement, &res.Engine.NoOfCylinders, &res.Engine.Range)

	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		_ = fmt.Errorf("unexpected error %v", err)
	}
}

// Delete is used for deleting values from database
func (d dB) Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	para := mux.Vars(r)
	Id := para["Id"]
	var value string
	err := d.DB.QueryRow("select Id from car where Id=?", Id).Scan(&value)

	if err != nil || value == "" {
		log.Printf("Unexpected error %v", err)
		w.Write([]byte("Entity do not exists"))
		return
	}
	d.DB.QueryRow("Delete  from car where Id=?", Id)
	d.DB.QueryRow("Delete  from engine where Car_id=?", Id)
	w.Write([]byte("Row deleted successfully"))
}

// Update is used for updating values in database.
func (d dB) Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	data := model.Car{}
	para := mux.Vars(r)
	Id := para["Id"]
	var value string
	err := d.DB.QueryRow("select Id from car where Id=?", Id).Scan(&value)
	if err != nil || value == "" {
		log.Printf("Unexpected error %v", err)
		w.Write([]byte("Entity do not exists"))
		return
	}
	res, _ := ioutil.ReadAll(r.Body)
	data.Id = uuid.MustParse(Id)
	err = json.Unmarshal(res, &data)
	if err != nil {
		_ = fmt.Errorf("unexpected error %v", err)
	}
	if data.Year < 1980 || data.Year > 2022 {
		w.Write([]byte("year must be between 1980 and 2022"))
		return
	}
	if data.Brand != "Tesla" && data.Brand != "Porsche" && data.Brand != "Ferrari" && data.Brand != "Mercedes" && data.Brand != "BMW" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if data.Fuel != "Petrol" && data.Fuel != "Diesel" && data.Fuel != "Electric" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if data.Fuel == "Electric" {
		data.Engine.Displacement = 0
		data.Engine.NoOfCylinders = 0
		if data.Engine.Range < 0 {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}
	if data.Fuel == "Petrol" || data.Fuel == "Diesel" {
		if data.Engine.Displacement <= 0 {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if data.Engine.NoOfCylinders <= 0 {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		data.Engine.Range = 0
	}
	_, err = d.DB.Exec("UPDATE car SET Name=? , Year=? , Brand=? , FuelType=? WHERE Id=?", data.Name, data.Year, data.Brand, data.Fuel, data.Id)
	if err != nil {
		_ = fmt.Errorf("unexpected error %v", err)
	}
	if err != nil {
		log.Printf("unexpected error %v", err)
	}
	_, err = d.DB.Exec("UPDATE engine SET Displacement=? , No_of_cylinders=? , `Range`=? WHERE Car_Id=?", data.Engine.Displacement, data.Engine.NoOfCylinders, data.Engine.Range, data.Id)
	if err != nil {
		log.Printf("unexpected error %v", err)
	}

	if err != nil {
		log.Printf("unexpected error %v", err)
	}
	json.NewEncoder(w).Encode(data)
}
