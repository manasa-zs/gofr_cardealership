package main

import (
	"developer.zopsmart.com/go/gofr/pkg/gofr"
	"hello-world/handler"
)

func main() {
	app := gofr.New()

	app.Server.ValidateHeaders = false
	app.EnableSwaggerUI()
	app.GET("/helloworld", handler.HelloWorld)
	app.Start()
}
