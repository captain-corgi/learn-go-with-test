package server

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/hotkratos/learn-go-with-test/cmd/tdd/server/stubs"
)

func TestHandler(t *testing.T) {
	data := "Hello, Anh"

	t.Run("1. Return data from store", func(t *testing.T) {
		store := stubs.NewSpyStore(data, t)
		sv := Server(store)

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()

		sv.ServeHTTP(res, req)

		got := res.Body.String()
		want := data
		if got != want {
			t.Errorf(`got "%s", want "%s"`, got, want)
		}
	})
	t.Run("2. Tells store to cancel work if request is cancelled", func(t *testing.T) {
		store := stubs.NewSpyStore(data, t)
		sv := Server(store)

		req := httptest.NewRequest(http.MethodGet, "/", nil)

		cancellingCtx, cancel := context.WithCancel(req.Context())
		time.AfterFunc(5*time.Millisecond, cancel)
		req = req.WithContext(cancellingCtx)

		res := stubs.NewSpyResponseWriter()

		sv.ServeHTTP(res, req)

		if res.Written() {
			t.Errorf("a response should not have been written")
		}
	})
}
