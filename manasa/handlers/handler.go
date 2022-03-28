package handlers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"manasa/models"
)

type Handler struct {
	car Service
}

// New function is a factory function for handler layer.
func New(car Service) Handler {
	return Handler{car: car}
}

// GetByBrand method is used in handler layer to fetch rows.
func (h Handler) GetByBrand(writer http.ResponseWriter, request *http.Request) {
	ctx := request.Context()

	brand := request.URL.Query().Get("brand")
	engine := request.URL.Query().Get("isEngine")

	isEngine, err := strconv.ParseBool(engine)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	res, err := h.car.GetByBrand(ctx, brand, isEngine)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)

		_, err = writer.Write([]byte(err.Error()))
		if err != nil {
			return
		}

		return
	}

	body, err := json.Marshal(res)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)

		_, err = writer.Write([]byte(err.Error()))
		if err != nil {
			return
		}

		return
	}

	writer.Header().Set("Content-Type", "application-json")

	_, err = writer.Write(body)
	if err != nil {
		log.Println("writing response err")
		return
	}
}

// GetByID method is used in handler layer to fetch row.
func (h Handler) GetByID(writer http.ResponseWriter, request *http.Request) {
	ctx := request.Context()

	writer.Header().Set("Content-Type", "application/json")

	param := mux.Vars(request)
	id := param["id"]

	ID, err := uuid.Parse(id)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)

		return
	}

	resp, err := h.car.GetByID(ctx, ID)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)

		_, err = writer.Write([]byte(err.Error()))
		if err != nil {
			return
		}

		return
	}

	body, err := json.Marshal(resp)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
	}

	_, err = writer.Write(body)
	if err != nil {
		log.Printf("%v", err)
	}
}

// Update method is used in handler layer to update rows.
func (h Handler) Update(writer http.ResponseWriter, request *http.Request) {
	ctx := request.Context()

	var car models.Car

	param := mux.Vars(request)
	paramID := param["id"]

	id, err := uuid.Parse(paramID)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	body, err := io.ReadAll(request.Body)
	if err != nil {
		return
	}

	err = json.Unmarshal(body, &car)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	car.ID = id

	res, err := h.car.Update(ctx, &car)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)

		_, err = writer.Write([]byte(err.Error()))
		if err != nil {
			return
		}

		return
	}

	body, err = json.Marshal(res)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-type", "application/json")

	writer.WriteHeader(http.StatusOK)

	_, err = writer.Write(body)
	if err != nil {
		log.Println("writing response error")
	}
}

// Create method is used in handler layer for inserting rows.
func (h Handler) Create(writer http.ResponseWriter, request *http.Request) {
	ctx := request.Context()

	writer.Header().Set("Content-Type", "application/json")

	var car models.Car

	body, err := io.ReadAll(request.Body)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(body, &car)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)

		_, err = writer.Write([]byte(err.Error()))
		if err != nil {
			return
		}

		return
	}

	resp, err := h.car.Create(ctx, &car)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)

		_, err = writer.Write([]byte(err.Error()))
		if err != nil {
			return
		}

		return
	}

	result, err := json.Marshal(resp)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusCreated)

	_, err = writer.Write(result)
	if err != nil {
		log.Println("writing response error")
		return
	}
}

// Delete method is used in handler layer for deleting rows.
func (h Handler) Delete(writer http.ResponseWriter, request *http.Request) {
	ctx := request.Context()

	writer.Header().Set("Content-Type", "application/json")

	param := mux.Vars(request)
	id := param["id"]

	ID, err := uuid.Parse(id)
	if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.car.Delete(ctx, ID)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusNoContent)
}
