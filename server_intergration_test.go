package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRecordingCommentsAndRetrievingThem(t *testing.T) {
	store := NewInMemoryUserStore()
	server := NewUserServer(store)
	user := "klim"

	server.ServeHTTP(httptest.NewRecorder(), newPostCommentRequest(user))
	server.ServeHTTP(httptest.NewRecorder(), newPostCommentRequest(user))
	server.ServeHTTP(httptest.NewRecorder(), newPostCommentRequest(user))

	response := httptest.NewRecorder()
	server.ServeHTTP(response, newGetPostsRequest(user))
	assertStatus(t, response.Code, http.StatusOK)

	assertResponseBody(t, response.Body.String(), "3")
}
