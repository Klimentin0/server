package server

import (
	"log"
	"net/http"
)

func main() {
	handler := http.HandlerFunc(UserServer)
	log.Fatal(http.ListenAndServe(":5000", handler))
}
