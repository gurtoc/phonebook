package main

import (
	"encoding/json"
	"fmt"
	mux2 "github.com/gorilla/mux"
	"log"
	"net/http"
)

type Person struct {
	ID          string `json:"id"`
	Firstname string `json:"firstname"`
	LastName string `json:"lastname"`
	City string `json:"city"`
	Zipcode int `json:"zipcode"`
	Number int `json:"number"`

}
//testest
//type
var phonebooks []Person

func singlephonebook(w http.ResponseWriter, r *http.Request){

	params := mux2.Vars(r)
	for _, p := range phonebooks {
		if p.ID == params["id"] {
			json.NewEncoder(w).Encode(p)
			return
		}
	}
	json.NewEncoder(w).Encode("Nie ma zadnych osob w ksiazce telefonicznej")
}

func overwievphonebook(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(phonebooks)
}

func postphonebook(w http.ResponseWriter, r *http.Request){
	var person Person
	_ = json.NewDecoder(r.Body).Decode(&person)
	phonebooks = append(phonebooks, person)
	json.NewEncoder(w).Encode(person)
	//fmt.Fprint(w, "testestest")
}

func homePage(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w,"Home PAge")
}


func handleRequest()  {

	mux := mux2.NewRouter().StrictSlash(true)

	mux.HandleFunc("/",homePage ).Methods("GET")
	mux.HandleFunc("/phonebook/",overwievphonebook ).Methods("GET")
	mux.HandleFunc("/phonebook/{id}",singlephonebook ).Methods("GET")
	mux.HandleFunc("/phonebook", postphonebook).Methods("POST")
	log.Fatal(http.ListenAndServe(":3456", mux))
}

func main() {
	phonebooks = append(phonebooks, Person{ID: "1", Firstname: "Jan", LastName: "Nowak", City: "Poznan", Zipcode: 62600, Number: 55455666})
	handleRequest()
}
