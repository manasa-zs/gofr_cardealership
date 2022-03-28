package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

type mockHandler struct{}

// ServeHTTP is a mock serveHttp method
func (m mockHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

// TestMiddleware checks if given api-key is valid
func TestMiddleware(t *testing.T) {
	testCases := []struct {
		desc       string
		value      string
		statusCode int
	}{
		{"Authentication Successful", "123456", http.StatusOK},
		{"Authentication Fail", "123", http.StatusUnauthorized},
	}

	for i, v := range testCases {
		req := httptest.NewRequest(http.MethodPost, "/car/create", nil)
		req.Header.Add("api-key", v.value)

		w := httptest.NewRecorder()
		a := Middleware(mockHandler{})

		a.ServeHTTP(w, req)
		res := w.Result()
		assert.Equal(t, v.statusCode, res.StatusCode, "Test case %d Failed: %s", i, v.desc)

		err := res.Body.Close()
		if err != nil {
			return
		}
	}
}
