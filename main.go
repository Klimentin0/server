package main

import (
	"log"
	"net/http"
)

func main() {
	server := NewUserServer(NewInMemoryUserStore())
	log.Fatal(http.ListenAndServe(":5000", server))
}
