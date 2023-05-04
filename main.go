package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	log.Println("server start at port 8081")
	log.Fatal(http.ListenAndServe(":8081", r))
}
