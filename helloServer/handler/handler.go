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
	name := query.Get("name")
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method not allowed"))

	} else if name == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Status bad request"))
	} else {
		w.Write([]byte("hello, " + name + "!"))
	}

}
