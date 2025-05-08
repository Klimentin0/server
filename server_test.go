package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

type StubUserStore struct {
	posts        map[string]int
	commentPosts []string
}

func TestGETusers(t *testing.T) {
	store := StubUserStore{
		map[string]int{
			"admin": 20,
			"klim":  10,
		},
		nil,
	}
	server := NewUserServer(&store)

	t.Run("returns admin's posts", func(t *testing.T) {
		request := newGetPostsRequest("admin")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusOK)
		assertResponseBody(t, response.Body.String(), "20")
	})

	t.Run("returns klim's posts", func(t *testing.T) {
		request := newGetPostsRequest("klim")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusOK)
		assertResponseBody(t, response.Body.String(), "10")
	})
	t.Run("returns 404 on missing users", func(t *testing.T) {
		request := newGetPostsRequest("Missing-Person")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusNotFound)
	})
}

func newGetPostsRequest(user string) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/users/%s", user), nil)
	return req
}

func assertStatus(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("did not get correct status, got %d, want %d", got, want)
	}
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

func (s *StubUserStore) PostComment(user string) {
	s.commentPosts = append(s.commentPosts, user)
}

func TestStoreComments(t *testing.T) {
	store := StubUserStore{
		map[string]int{},
		nil,
	}
	server := NewUserServer(&store)

	t.Run("it returns accepted on POST", func(t *testing.T) {
		user := "klim"
		request := newPostCommentRequest(user)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusAccepted)

		if len(store.commentPosts) != 1 {
			t.Fatalf("got %d calls to PostComment want %d", len(store.commentPosts), 1)
		}

		if store.commentPosts[0] != user {
			t.Errorf("did not store correct winner got %q want %q", store.commentPosts[0], user)
		}
	})
}

func newPostCommentRequest(user string) *http.Request {
	req, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("/users/%s", user), nil)
	return req
}

func TestBlog(t *testing.T) {
	store := StubUserStore{}
	server := NewUserServer(&store)

	t.Run("it returns 200 on /blog", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/blog", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusOK)
	})
}
