package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Person struct de Persona
type Person struct {
	ID        string   `json:"id,omitempty`
	FirstName string   `json:"firstname,omitempty"`
	LastName  string   `json:"lastname,omitempty"`
	Address   *Address `json:"address,omitempty"`
}

// Address struct de Direccion, agregando Ciudad y Comuna
type Address struct {
	City    string `json:"city,omitempty"`
	Commune string `json:"commune,omitempty"`
}

//Arreglo de Persona llamado people
var people []Person

// GetPeopleEndpoint obtiene todas las Personas del arreglo
func GetPeopleEndpoint(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(people)

}

// GetPersonEndpoint obtiene una persona por ID en donde /people/getPerson/{id}
func GetPersonEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for _, item := range people {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}

	}
	json.NewEncoder(w).Encode(&Person{})
}

// CreatePersonEndpoint creamos a la persona por ID en donde /people/createPerson/{id}
func CreatePersonEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	var person Person
	_ = json.NewDecoder(req.Body).Decode(&person)
	person.ID = params["id"]
	people = append(people, person)
	json.NewEncoder(w).Encode(people)
}

// DeletePersonEndpoint Elimina una persona por ID en donde /people/deletePerson/{id}
func DeletePersonEndpoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for index, item := range people {
		if item.ID == params["id"] {
			people = append(people[:index], people[index+1:]...)
			break
		}

	}
	json.NewEncoder(w).Encode(people)
}

//main
func main() {

	router := mux.NewRouter()
	people = append(people, Person{ID: "1", FirstName: "Paris", LastName: "Aguilera", Address: &Address{City: "Santiago", Commune: "Lampa"}})
	people = append(people, Person{ID: "2", FirstName: "Bastian", LastName: "Moreno"})
	people = append(people, Person{ID: "3", FirstName: "Sofia", LastName: "Uppa", Address: &Address{City: "Santiago", Commune: "La Florida"}})

	router.HandleFunc("/people", GetPeopleEndpoint).Methods("GET")
	router.HandleFunc("/people/getPerson/{id}", GetPersonEndpoint).Methods("GET")
	router.HandleFunc("/people/createPerson/{id}", CreatePersonEndpoint).Methods("POST")
	router.HandleFunc("/people/deletePerson/{id}", DeletePersonEndpoint).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":3000", router))
}
