package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"log"
	"manasa/handlers"
	"manasa/middleware"
	"manasa/services"
	"manasa/stores/car"
	"manasa/stores/engine"
	"net/http"
)

// DBConn function is used for connecting with DB.
func DBConn() (db *sql.DB, err error) {
	db, err = sql.Open("mysql", "manasa:Manasa@2210@tcp(docker.linux.for.localhost:3306)/Dealership")
	if err != nil {
		return nil, err
	}

	return db, nil
}

func main() {
	db, err := DBConn()
	if err != nil {
		log.Fatalf("cannot connect to datastore %v", err)
	}

	log.Println("Successfully connected to db")

	carStore := car.New(db)
	engineStore := engine.New(db)
	service := services.New(carStore, engineStore)
	h := handlers.New(service)

	r := mux.NewRouter()
	r.Use(middleware.Middleware)
	r.HandleFunc("/cars", h.GetByBrand).Methods(http.MethodGet)
	r.HandleFunc("/cars/{id}", h.GetByID).Methods(http.MethodGet)
	r.HandleFunc("/cars/{id}", h.Update).Methods(http.MethodPut)
	r.HandleFunc("/cars", h.Create).Methods(http.MethodPost)
	r.HandleFunc("/cars/{id}", h.Delete).Methods(http.MethodDelete)

	err = http.ListenAndServe(":8000", r)
	if err != nil {
		log.Println("error occurred in TCP server")
	}
}
