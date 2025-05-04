package server

import (
	"fmt"
	"net/http"
	"strings"
)

func UserServer(w http.ResponseWriter, r *http.Request) {
	user := strings.TrimPrefix(r.URL.Path, "/users/")

	if user == "admin" {
		fmt.Fprint(w, "20")
		return
	}

	if user == "klim" {
		fmt.Fprint(w, "10")
		return
	}
}
