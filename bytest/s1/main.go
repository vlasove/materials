package main

import (
	"log"
	"net/http"
)

// Todo: add context
// Todo: add PostgresPlayerStore
// Todo: add CI with github-actions
func main() {
	server := &PlayerServer{NewInMemoryPlayerStore()}
	log.Fatal(http.ListenAndServe(":8080", server))
}
