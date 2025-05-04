package main

import (
	"log"
	"net/http"
)

type InMemoryUserStore struct{}

func (i *InMemoryUserStore) GetUserPosts(user string) int {
	return 123
}

func main() {
	server := &UserServer{&InMemoryUserStore{}}
	log.Fatal(http.ListenAndServe(":5000", server))
}
