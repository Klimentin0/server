package main

import (
	"fmt"
	"net/http"
	"strings"
)

type UserStore interface {
	GetUserPosts(user string) int
}

type UserServer struct {
	store UserStore
}

func (p *UserServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	user := strings.TrimPrefix(r.URL.Path, "/users/")

	fmt.Fprint(w, p.store.GetUserPosts(user))
}
