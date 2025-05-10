package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type User struct {
	Name     string
	Comments int
}

type UserStore interface {
	GetUserPosts(user string) int
	PostComment(user string)
	GetBlog() []User
}

type UserServer struct {
	store UserStore
	http.Handler
}

const jsonContentType = "application/json"

func NewUserServer(store UserStore) *UserServer {
	p := new(UserServer)
	p.store = store
	router := http.NewServeMux()
	router.Handle("/blog", http.HandlerFunc(p.blogHandler))
	router.Handle("/users/", http.HandlerFunc(p.usersHandler))

	p.Handler = router

	return p
}

func (p *UserServer) blogHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", jsonContentType)
	json.NewEncoder(w).Encode(p.store.GetBlog())
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
