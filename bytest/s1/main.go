package main

import (
	"log"
	"net/http"
)

// Todo: add PostgresPlayerStore
// Todo: add context

func main() {
	server := &PlayerServer{NewInMemoryPlayerStore()}
	log.Fatal(http.ListenAndServe(":8080", server))
}
