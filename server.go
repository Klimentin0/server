package main

import (
	"fmt"
	"net/http"
	"strings"
)

type UserStore interface {
	GetUserPosts(user string) int
	PostComment(user string)
}

type UserServer struct {
	store UserStore
	http.Handler
}

func NewUserServer(store UserStore) *UserServer {
	p := new(UserServer)
	p.store = store
	router := http.NewServeMux()
	router.Handle("/blog", http.HandlerFunc(p.blogHandler))
	router.Handle("/users", http.HandlerFunc(p.usersHandler))

	p.Handler = router

	return p
}

func (p *UserServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	p.ServeHTTP(w, r)
}

func (p *UserServer) blogHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (p *UserServer) usersHandler(w http.ResponseWriter, r *http.Request) {
	user := strings.TrimPrefix(r.URL.Path, "/users/")
	switch r.Method {
	case http.MethodPost:
		p.processComments(w, user)
	case http.MethodGet:
		p.showPosts(w, user)
	}
}

func (p *UserServer) showPosts(w http.ResponseWriter, user string) {
	posts := p.store.GetUserPosts(user)

	if posts == 0 {
		w.WriteHeader(http.StatusNotFound)
	}
	fmt.Fprint(w, posts)
}

func (p *UserServer) processComments(w http.ResponseWriter, user string) {
	p.store.PostComment(user)
	w.WriteHeader(http.StatusAccepted)
}
