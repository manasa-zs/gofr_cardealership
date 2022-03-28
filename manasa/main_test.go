package main

import (
	"bytes"
	"fmt"
	"net/http"
	"testing"
)

func Test_Main(t *testing.T) {
	testcases := []struct {
		desc   string
		method string
		url    string
		body   []byte
		status int
	}{
		{"success case create", http.MethodPost, "http://localhost:8000/cars",
			[]byte(`{"name": "bmw 2", "year": 2000,"brand":"tesla","fuel":"electric","engine":{"range":40}}`), http.StatusCreated},
		{"Success case GetById", http.MethodGet, "http://localhost:8000/cars/22148aa7-49a0-4b6d-8078-11baa025cf4a", nil, http.StatusOK},
		{"Success case GetByBrand", http.MethodGet, "http://localhost:8000/cars?brand=BMW&&isEngine=true", nil, http.StatusOK},
		{"Success case Update", http.MethodPut, "http://localhost:8000/cars/22148aa7-49a0-4b6d-8078-11baa025cf4a",
			[]byte(`{"name": "bmw 2", "year": 2001,"brand":"tesla","fuel":"electric","engine":{"range":40}}`), http.StatusOK},
		{"invalid parameters to create", http.MethodPost, "http://localhost:8000/cars",
			[]byte(`{"name": "bmw 2", "year": 2000,"brand":"tesla","fuel":"electric","engine":{"range":40}`), http.StatusBadRequest},
		{"invalid parameters to getById", http.MethodGet, "http://localhost:8000/cars/ea3", nil, http.StatusBadRequest},
		{"invalid parameters to getByBrand", http.MethodGet, "http://localhost:8000/cars?brand=BMW&&isEngine=tr", nil, http.StatusBadRequest},
		{"invalid parameters to Update", http.MethodPut, "http://localhost:8000/cars/ea3",
			[]byte(`{"name": "bmw 2", "year": 2001,"brand":"tesla","fuel":"electric","engine":{"range":40}}`), http.StatusBadRequest},
		{"invalid parameter to Delete", http.MethodDelete, "http://localhost:8000/cars/2da6f147c2", nil, http.StatusBadRequest},
		{"invalid year for create", http.MethodPost, "http://localhost:8000/cars",
			[]byte(`{"name": "bmw 2", "year": 1678,"brand":"tesla","fuel":"electric","engine":{"range":40}}`), http.StatusInternalServerError},
		{"invalid brand for create", http.MethodPost, "http://localhost:8000/cars",
			[]byte(`{"name": "bmw 2", "year": 2010,"brand":"maruthi","fuel":"electric","engine":{"range":40}}`), http.StatusInternalServerError},
		{"invalid engine details for create", http.MethodPost, "http://localhost:8000/cars",
			[]byte(`{"name": "bmw 2", "year": 1678,"brand":"tesla","fuel":"petrol","engine":{"range":40}}`), http.StatusInternalServerError},
		{"id doesn't exists GetById", http.MethodGet, "http://localhost:8000/cars/ea362c02-b998-46ec-907c-84aca7e858df",
			nil, http.StatusInternalServerError},
		{"id doesn't exists Update", http.MethodPut, "http://localhost:8000/cars/ea362c02-b998-46ec-907c-84aca7e858df",
			[]byte(`{"name": "bmw 2", "year": 2001,"brand":"tesla","fuel":"electric","engine":{"range":40}}`), http.StatusInternalServerError},
		{"id doesn't Delete", http.MethodDelete, "http://localhost:8000/cars/fc351219-675b-443c-89c1-222da6f148d2",
			nil, http.StatusInternalServerError},
		//{"Success case Delete", http.MethodDelete, "http://localhost:8000/cars/00ec8552-d676-4434-ac92-ec27c81772f0", nil, http.StatusNoContent},
	}

	for _, tc := range testcases {
		req, err := http.NewRequest(tc.method, tc.url, bytes.NewBuffer(tc.body))
		if err != nil {
			fmt.Println(err)
		}

		req.Header.Set("api-key", "123456")

		h := http.Client{}

		res, err := h.Do(req)
		if err != nil {
			fmt.Println(err)
		}

		res.Body.Close()

		if tc.status != res.StatusCode {
			t.Errorf("In method %v expected status %v got %v\n", tc.desc, tc.status, res.StatusCode)
		}
	}
}
