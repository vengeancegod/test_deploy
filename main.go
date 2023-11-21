package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Person struct {
	Name string
	Age  int
}

var people []Person

func main() {
	http.HandleFunc("/people", peopleHandler)
	http.HandleFunc("/condition", conditionCheckHandler)
	http.HandleFunc("/", dockerHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func peopleHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getPeople(w, r)
	case http.MethodPost:
		postPerson(w, r)
	default:
		http.Error(w, "invalid http", http.StatusMethodNotAllowed)
	}
}

func getPeople(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(people)
	fmt.Fprintf(w, "add new person: '%v'", people)
}

func postPerson(w http.ResponseWriter, r *http.Request) {
	var person Person
	err := json.NewDecoder(r.Body).Decode(&person)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	people = append(people, person)
	fmt.Fprintf(w, "add new people: '%v'", person)
}

func dockerHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello from docker")
}

func conditionCheckHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "server work correctly")
}
