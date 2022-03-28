package main

import (
	"helloServer/handler"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", handler.Handler)
	http.ListenAndServe(":8000", nil)
}
