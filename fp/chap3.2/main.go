package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

const (
	port = ":8080"
)

func init() {
	log.SetFlags(log.Lshortfile | log.Ldate)
	log.SetOutput(os.Stdout)
}

func main() {
	if os.Getenv("RUN_HTTP_SERVER") == "TRUE" {
		r := mux.NewRouter()
		r.HandleFunc("/cars", hof.CarsHandler).Methods("GET")
		r.HandleFunc("/cars/{id}", hof.CardsIdHandler).Methods("GET")

		log.Println("server starting at port", port)
		if err := http.ListenAndServe(port, r); err != nil {
			log.Fatal(err)
		}
	}
}
