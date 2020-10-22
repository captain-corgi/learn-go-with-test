package usecase

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/hotkratos/learn-go-with-test/internal/svc/players/usecase/stubs"
)

func TestGETPlayer(t *testing.T) {
	stubPlayerRepo := stubs.NewStubPlayerStore(
		map[string]int{
			"Anh":   20,
			"AnhLe": 40,
		},
	)
	server := NewPlayerServer(stubPlayerRepo)

	t.Run("returns Anh's score", func(t *testing.T) {
		req := getScoreRequest("Anh")
		res := httptest.NewRecorder()

		server.ServeHTTP(res, req)

		assertResponseBody(t, res.Body.String(), 20)
	})

	t.Run("returns AnhLe's score", func(t *testing.T) {
		req := getScoreRequest("AnhLe")
		res := httptest.NewRecorder()

		server.ServeHTTP(res, req)

		assertResponseBody(t, res.Body.String(), 40)
	})

	t.Run("returns 404 if user not found", func(t *testing.T) {
		req := getScoreRequest("LeTuanAnh")
		res := httptest.NewRecorder()

		server.ServeHTTP(res, req)

		got := res.Code
		want := 404

		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
}

func getScoreRequest(name string) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s", name), nil)
	return req
}

func assertResponseBody(t *testing.T, got string, want int) {
	t.Helper()
	if got != strconv.Itoa(want) {
		t.Errorf("got %q, want %q", got, want)
	}
}
