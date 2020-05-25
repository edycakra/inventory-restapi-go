package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homePage).Methods("GET")
	log.Fatal(http.ListenAndServe(":3000", router))
}

func main() {}
