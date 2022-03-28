package main

import (
	"cardealership/handler"
	"cardealership/services"
	"cardealership/stores/car"
	"cardealership/stores/engine"
	"developer.zopsmart.com/go/gofr/pkg/gofr"
)

func main() {
	app := gofr.New()

	carStore := car.New()
	engineStore := engine.New()
	service := services.New(carStore, engineStore)
	handler := handler.New(service)

	app.Server.ValidateHeaders = false
	app.POST("/cars", handler.Create)
	app.GET("/cars/{id}", handler.GetByID)
	app.GET("/cars", handler.GetByBrand)
	app.PUT("/cars/{id}", handler.Update)
	app.DELETE("/cars/{id}", handler.Delete)
	app.Start()
}
