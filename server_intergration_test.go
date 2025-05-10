package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRecordingCommentsAndRetrievingThem(t *testing.T) {
	store := NewInMemoryUserStore()
	server := NewUserServer(store)
	user := "Kyle"

	for i := 0; i < 42; i++ {
		server.ServeHTTP(httptest.NewRecorder(), newPostCommentRequest(user))
	}

	t.Run("get comments", func(t *testing.T) {
		response := httptest.NewRecorder()
		server.ServeHTTP(response, newGetPostsRequest(user))
		assertStatus(t, response.Code, http.StatusOK)

		assertResponseBody(t, response.Body.String(), "42")
	})
	t.Run("get blog", func(t *testing.T) {
		response := httptest.NewRecorder()
		server.ServeHTTP(response, newBlogRequest())
		assertStatus(t, response.Code, http.StatusOK)

		got := getBlogFromResponse(t, response.Body)
		want := []User{
			{"Kyle", 42},
		}
		assertBlog(t, got, want)
	})
}
