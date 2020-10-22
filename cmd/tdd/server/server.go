package server

import (
	"fmt"
	"net/http"
)

//Server is a simple store http handler function
func Server(store Store) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		data, err := store.Fetch(r.Context())

		if err != nil {
			return
		}

		fmt.Fprint(rw, data)
	}
}
