package middleware

import (
	"net/http"
)

// Middleware function is used for authentication.
func Middleware(handlerFunc http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		if request.Header.Get("api-key") != "123456" {
			writer.WriteHeader(http.StatusUnauthorized)
			return
		}

		handlerFunc.ServeHTTP(writer, request)
	})
}
