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
	contactInfo `json:"contactInfo"`

}

type contactInfo struct {
	City string `json:"city"`
	Zipcode int `json:"zipcode"`
	Number int `json:"number""`
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

func deletephonebook(w http.ResponseWriter, r *http.Request) {
	params := mux2.Vars(r)
	for i, p := range phonebooks {
		if p.ID == params["id"] {
			copy(phonebooks[i:], phonebooks[i+1:])
			phonebooks = phonebooks[:len(phonebooks)-1]
			break
		}
	}
	json.NewEncoder(w).Encode(phonebooks)
}


func updatephonebook(w http.ResponseWriter, r *http.Request) {
	var person Person
	_ = json.NewDecoder(r.Body).Decode(&person)
	params := mux2.Vars(r)
	for i, p := range phonebooks {
		if p.ID == params["id"] {
			phonebooks[i] = person
			json.NewEncoder(w).Encode(person)
			break
		}
	}
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
	mux.HandleFunc("/phonebook/{id}", deletephonebook).Methods("DELETE")
	mux.HandleFunc("/phonebook/{id}",updatephonebook ).Methods("PUT")
	log.Fatal(http.ListenAndServe(":3456", mux))
	//fmt.Println("Server działa na localport:3456")
}



func main() {
	//phonebooks = append(phonebooks, Person{ID: "1", Firstname: "Jan", LastName: "Nowak", City: "Poznan", Zipcode: 62600, Number: 55455666})
	phonebooks = append(phonebooks, Person{ID: "1", Firstname: "Jan", LastName: "Nowak",contactInfo: contactInfo{ City: "Poznan", Zipcode: 62600, Number: 55455666}})
	phonebooks = append(phonebooks, Person{ID: "2", Firstname: "Franek", LastName: "Kimono",contactInfo: contactInfo{ City: "Warszawa", Zipcode: 87542, Number: 997}})

	handleRequest()
}
