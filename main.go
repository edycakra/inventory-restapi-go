package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
)

//Item for inventories
type Item struct { 
	UID string `json:"UID"`
	Name string `json:"Name"` 
	Desc string `json:"Desc"`
	Price float64 `json:"Price"`
}

var inventory []Item

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Endpoint homePage is called")
}

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homePage).Methods("GET")
	log.Fatal(http.ListenAndServe(":3000", router))
}

func main() {
	handleRequests()
}
