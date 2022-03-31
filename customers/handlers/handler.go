package handlers

import (
	"developer.zopsmart.com/go/gofr/pkg/errors"
	"developer.zopsmart.com/go/gofr/pkg/gofr"
	"developer.zopsmart.com/go/gofr/pkg/gofr/types"
	"github.com/google/uuid"
	"training/customers/models"
)

type handler struct {
	service CustomerService
}

func New(h CustomerService) *handler {
	return &handler{service: h}
}

func (h *handler) Create(ctx *gofr.Context) (interface{}, error) {
	var customer *models.Customer
	var respCustomer *models.Customer
	err := ctx.Bind(&customer)
	if err != nil {
		return nil, errors.InvalidParam{Param: []string{"Body"}}
	}

	respCustomer, err = h.service.Create(ctx, customer)
	if err != nil {
		return nil, err
	}
	resp := types.Response{
		Data: respCustomer,
	}

	return resp, nil
}
func (h *handler) Get(ctx *gofr.Context) (interface{}, error) {
	var customer *models.Customer
	param := ctx.PathParam("id")
	ID, err := uuid.Parse(param)
	if err != nil {
		return nil, err
	}

	customer, err = h.service.Get(ctx, ID)
	if err != nil {
		return nil, err
	}

	return customer, nil
}

func (h *handler) Update(ctx *gofr.Context) (interface{}, error) {
	var customer *models.Customer
	var respCustomer *models.Customer
	param := ctx.PathParam("id")
	ID, err := uuid.Parse(param)
	if err != nil {
		return nil, err
	}

	err = ctx.Bind(&customer)
	if err != nil {
		return nil, errors.InvalidParam{Param: []string{"Body"}}
	}

	customer.ID = ID
	respCustomer, err = h.service.Update(ctx, customer)
	if err != nil {
		return nil, err
	}
	resp := types.Response{
		Data: respCustomer,
	}

	return resp, nil
}

func (h *handler) Delete(ctx *gofr.Context) (interface{}, error) {
	param := ctx.PathParam("id")
	ID, err := uuid.Parse(param)
	if err != nil {
		return nil, err
	}

	err = h.service.Delete(ctx, ID)
	if err != nil {
		return nil, err
	}

	return nil, nil
}
