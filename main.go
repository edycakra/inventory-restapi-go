package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//Item for inventories
type Item struct {
	UID   string  `json:"UID"`
	Name  string  `json:"Name"`
	Desc  string  `json:"Desc"`
	Price float64 `json:"Price"`
}

var inventory []Item

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Endpoint homePage is called")
}

func getItems(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("Calling getItems")
	json.NewEncoder(w).Encode(inventory)
}

func createItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var item Item
	_ = json.NewDecoder(r.Body).Decode(&item) //get the data that we input

	inventory = append(inventory, item) //store the data

	json.NewEncoder(w).Encode(item)
}

func deleteItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	_deleteItemUID(params["uid"])

	json.NewEncoder(w).Encode(inventory)
}

func _deleteItemUID(uid string) {
	for index, item := range inventory {
		if item.UID == uid {
			//delete by slicing
			inventory = append(inventory[:index], inventory[index+1:]...)
			break
		}
	}
}

func updateItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var item Item
	_ = json.NewDecoder(r.Body).Decode(&item) //get the data that we input
	params := mux.Vars(r)
	_deleteItemUID(params["uid"])
	inventory = append(inventory, item) //store the data
	json.NewEncoder(w).Encode(inventory)

}

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)

	//list of routes
	router.HandleFunc("/", homePage).Methods("GET")
	router.HandleFunc("/inventory", getItems).Methods("GET")
	router.HandleFunc("/inventory", createItem).Methods("POST")
	router.HandleFunc("/inventory/{uid}", deleteItem).Methods("DELETE")
	router.HandleFunc("/inventory/{uid}", updateItem).Methods("PUT")

	log.Fatal(http.ListenAndServe(":3000", router))
}

func main() {
	inventory = append(inventory, Item{
		UID:   "0",
		Name:  "Cheese",
		Desc:  "Pasta Filata",
		Price: 4.99,
	})
	inventory = append(inventory, Item{
		UID:   "1",
		Name:  "Milk",
		Desc:  "Lowfat Milk",
		Price: 3.25,
	})
	handleRequests()
}
