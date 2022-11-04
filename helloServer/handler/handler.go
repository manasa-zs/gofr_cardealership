package handler

import (
	"net/http"
)

// Handler used to send request to the server 8000.
func Handler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	if len(query) == 0 && r.Method == http.MethodGet {
		w.Write([]byte("hello"))
		return
	}
	code := query.Get("code")
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method not allowed"))
	} else if code == "400" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Status bad request"))
	} else if code == "401" {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Status unauthorized"))
	} else if code == "403" {
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte("Status forbidden"))
	} else if code == "404" {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Status bad request"))
	} else if code == "500" {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Status Internal Server Error"))
	} else {
		w.Write([]byte("hello, " + code + "!"))
	}

}
