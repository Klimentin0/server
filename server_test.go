package server

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGETusers(t *testing.T) {
	t.Run("returns admin's posts", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/users/admin", nil)
		response := httptest.NewRecorder()

		UserServer(response, request)

		got := response.Body.String()
		want := "20"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("returns klim's posts", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/users/klim", nil)
		response := httptest.NewRecorder()

		UserServer(response, request)

		got := response.Body.String()
		want := "10"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}
