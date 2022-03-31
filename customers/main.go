package main

import (
	"developer.zopsmart.com/go/gofr/pkg/gofr"
	"training/customers/handlers"
	"training/customers/services"
	"training/customers/stores"
)

func main() {
	app := gofr.New()
	app.Server.ValidateHeaders = false

	store := stores.New()
	service := services.New(store)
	handler := handlers.New(service)
	app.POST("/customers", handler.Create)
	app.GET("/customers/{id}", handler.Get)
	app.PUT("/customers/{id}", handler.Update)
	app.DELETE("/customers/{id}", handler.Delete)
	app.Start()
}
