package main

import (
	"encoding/json"
	"fmt"
	mux2 "github.com/gorilla/mux"
	"log"
	"net/http"
)

type phonebook struct {
	Firstname string `json:"firstname"`
	LastName string `json:"lastname"`
	City string `json:"city"`
	Zipcode int `json:"zipcode"`
	Number int `json:"number"`

}
//testest
type phonebooks []phonebook

func allphonebook(w http.ResponseWriter, r *http.Request){
	phonebooks := phonebooks{
		phonebook{"Imie", "Nazwisko", "Miasto", 62 - 600, 1234567},
	}
	fmt.Println("test")
	json.NewEncoder(w).Encode(phonebooks)
}

func postphonebook(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, "testestest")
}

func homePage(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w,"Home PAge")
}


func handleRequest()  {
	mux := mux2.NewRouter().StrictSlash(true)

	mux.HandleFunc("/", homePage)
	mux.HandleFunc("/phonebook", allphonebook).Methods("GET")
	mux.HandleFunc("/phonebook", postphonebook).Methods("POST")
	log.Fatal(http.ListenAndServe(":8083", mux))
}

func main() {
	handleRequest()

}