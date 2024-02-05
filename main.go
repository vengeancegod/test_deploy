package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-redis/redis/v8"
)

// asdasdas
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

var ctx = context.Background()

func dockerHandler(w http.ResponseWriter, r *http.Request) {
	bd := redis.NewClient(&redis.Options{
		Addr: "amvera-darkyz-run-test-deploy2:6379",
		//Password: "eqgadya6YZAKY9alFQmDJzAKnfaQYL750WayS3HFT9k",
		DB: 0,
	})
	_, err := bd.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Ошибка подключения: %v", err)
	}
	if r.Method == http.MethodGet && r.URL.Path == "/" {
		err := bd.Incr(ctx, "counter").Err()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		counter, err := bd.Get(ctx, "counter").Result()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		fmt.Fprintf(w, "Счетчик: %s", counter)
		defer bd.Close()
	}
}

func conditionCheckHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "server work correctly!!!")
}
