package transport

import "net/http"

type (
	//HTTPHandler is an http handler
	HTTPHandler interface {
		ServeHTTP(rw http.ResponseWriter, req *http.Request)
	}
)
