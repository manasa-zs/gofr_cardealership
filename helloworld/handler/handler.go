package handler

import "developer.zopsmart.com/go/gofr/pkg/gofr"

func HelloWorld(ctx *gofr.Context) (interface{}, error) {
	return "HelloWorld", nil
}
