package stubs

import (
	"errors"
	"net/http"
)

//SpyResponseWriter is a response writer spy
type SpyResponseWriter struct {
	written bool
}

//NewSpyResponseWriter return a new SpyResponseWriter pointer, written is false
func NewSpyResponseWriter() *SpyResponseWriter {
	return &SpyResponseWriter{false}
}

//Header set written = true
func (r *SpyResponseWriter) Header() http.Header {
	r.written = true
	return nil
}

func (r *SpyResponseWriter) Write([]byte) (int, error) {
	r.written = true
	return 0, errors.New("not implemented")
}

//WriteHeader set written = true
func (r *SpyResponseWriter) WriteHeader(statusCode int) {
	r.written = true
}

//Written return value of written
func (r *SpyResponseWriter) Written() bool {
	return r.written
}
