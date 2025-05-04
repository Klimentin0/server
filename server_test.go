package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

type StubUserStore struct {
	posts map[string]int
}

func TestGETusers(t *testing.T) {
	store := StubUserStore{
		map[string]int{
			"admin": 20,
			"klim":  10,
		},
	}
	server := &UserServer{&store}

	t.Run("returns admin's posts", func(t *testing.T) {
		request := newGetPostsRequest("admin")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertResponseBody(t, response.Body.String(), "20")
	})

	t.Run("returns klim's posts", func(t *testing.T) {
		request := newGetPostsRequest("klim")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertResponseBody(t, response.Body.String(), "10")
	})
}

func newGetPostsRequest(user string) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/users/%s", user), nil)
	return req
}

func assertResponseBody(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("response body is wrong, got %q watn %q", got, want)
	}
}

func (s *StubUserStore) GetUserPosts(user string) int {
	posts := s.posts[user]
	return posts
}
