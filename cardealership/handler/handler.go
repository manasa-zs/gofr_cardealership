package handler

import (
	"cardealership/models"
	"developer.zopsmart.com/go/gofr/pkg/errors"
	"developer.zopsmart.com/go/gofr/pkg/gofr"
	"developer.zopsmart.com/go/gofr/pkg/gofr/types"
	"github.com/google/uuid"
	"strconv"
)

type response struct {
	Data interface{}
}
type APIHandler struct {
	serviceHandler Service
}

func New(car Service) *APIHandler {
	return &APIHandler{serviceHandler: car}
}

// Create method is used in handler layer for inserting rows.
func (h *APIHandler) Create(ctx *gofr.Context) (interface{}, error) {

	var car *models.Car
	var respCar *models.Car
	err := ctx.Bind(&car)
	if err != nil {
		return nil, errors.InvalidParam{Param: []string{"Body"}}
	}

	respCar, err = h.serviceHandler.Create(ctx, car)
	if err != nil {
		return nil, err
	}
	resp := types.Response{
		Data: response{
			Data: respCar,
		},
	}

	return resp, nil
}

func (h *APIHandler) Delete(ctx *gofr.Context) (interface{}, error) {
	id := ctx.PathParam("id")
	ID, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	err = h.serviceHandler.Delete(ctx, ID)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (h *APIHandler) Update(ctx *gofr.Context) (interface{}, error) {
	var car *models.Car
	var respCar *models.Car

	param := ctx.PathParam("id")

	id, err := uuid.Parse(param)
	if err != nil {
		return nil, err
	}

	err = ctx.Bind(&car)
	if err != nil {
		return nil, errors.InvalidParam{Param: []string{"Body"}}
	}

	car.ID = id

	respCar, err = h.serviceHandler.Update(ctx, car)
	if err != nil {
		return nil, err
	}
	resp := types.Response{
		Data: response{
			Data: respCar,
		},
	}

	return resp, nil
}

func (h *APIHandler) GetByID(ctx *gofr.Context) (interface{}, error) {
	id := ctx.PathParam("id")

	ID, err := uuid.Parse(id)
	if err != nil {
		return nil, errors.InvalidParam{Param: []string{"id"}}
	}

	resp, err := h.serviceHandler.GetByID(ctx, ID)

	if err == nil {
		resp := types.Response{
			Data: response{
				Data: resp,
			},
		}
		return resp, nil
	}

	return nil, errors.Error("err")
}

func (h *APIHandler) GetByBrand(ctx *gofr.Context) (interface{}, error) {
	brand := ctx.Param("brand")
	engine := ctx.Param("isEngine")

	isEngine, err := strconv.ParseBool(engine)
	if err != nil {
		return nil, errors.InvalidParam{Param: []string{"isEngine"}}
	}

	res, err := h.serviceHandler.GetByBrand(ctx, brand, isEngine)
	if err != nil {
		return nil, err
	}
	return res, nil
}
