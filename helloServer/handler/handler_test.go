package handler

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestHelloServer is used the testcases for
func TestHelloServer(t *testing.T) {
	Testcases := []struct {
		des      string
		method   string
		input    string
		expected string
	}{
		{"test case with query parameter", http.MethodGet, "/hello?name=manu", "hello, manu!"},
		{"test case without query parameter", http.MethodGet, "/hello", "hello"},
		{"test case without query parameter", http.MethodGet, "/hello?name=", "Status bad request"},
		{"test case with method other than Get", http.MethodPost, "/hello?name=manu", "Method not allowed"},
		{"test case with method other than Get", http.MethodPost, "/hello", "Method not allowed"},
		{"test case with method other than Get", http.MethodPut, "/hello?name=manu", "Method not allowed"},
		{"test case with method other than Get", http.MethodPut, "/hello", "Method not allowed"},
		{"test case with method other than Get", http.MethodPatch, "/hello?name=manu", "Method not allowed"},
		{"test case with method other than Get", http.MethodPatch, "/hello", "Method not allowed"},
		{"test case with method other than Get", http.MethodDelete, "/hello?name=manu", "Method not allowed"},
		{"test case with method other than Get", http.MethodDelete, "/hello", "Method not allowed"},
	}
	for _, v := range Testcases {
		req := httptest.NewRequest(v.method, v.input, nil)
		w := httptest.NewRecorder()
		Handler(w, req)
		res := w.Result()
		defer res.Body.Close()

		data, err := ioutil.ReadAll(res.Body)
		if err != nil {
			t.Errorf("expected error to be nil got %v", err)
		}

		if string(data) != v.expected {
			t.Errorf("expected %v got %v", v.expected, string(data))
		}
	}
}

func BenchmarkHandler(b *testing.B) {

	for i := 0; i < b.N; i++ {
		req := httptest.NewRequest(http.MethodGet, "/hello?name=manu", nil)
		w := httptest.NewRecorder()

		Handler(w, req)
	}
}
