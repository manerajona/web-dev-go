package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Message struct {
	Id     string       `json:"id"`
	Title  string       `json:"title"`
	Body   string       `json:"body"`
	TypeMs *TypeMessage `json:"typeMessage"`
}

type TypeMessage struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

var messages []Message

func findAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(messages)
}

func getMessage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range messages {
		if item.Id == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Message{})
}

func create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var message Message
	json.NewDecoder(r.Body).Decode(&message)
	message.Id = strconv.Itoa(rand.Intn(10000000))
	messages = append(messages, message)
	json.NewEncoder(w).Encode(message)
}

func update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range messages {
		if item.Id == params["id"] {
			messages = append(messages[:index], messages[index+1:]...)
			var message Message
			json.NewDecoder(r.Body).Decode(&message)
			messages = append(messages, message)
			json.NewEncoder(w).Encode(message)
			return
		}
	}
}

func delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range messages {
		if item.Id == params["id"] {
			messages = append(messages[:index], messages[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(messages)
}

func main() {
	r := mux.NewRouter()

	// Mock data
	messages = append(messages, Message{Id: "1", Title: "hello", Body: "Hello world!", TypeMs: &TypeMessage{Id: "1", Name: "Greeting"}})
	messages = append(messages, Message{Id: "2", Title: "bye", Body: "Bye world!", TypeMs: &TypeMessage{Id: "1", Name: "Greeting"}})
	messages = append(messages, Message{Id: "3", Title: "Go", Body: "Go is cool", TypeMs: &TypeMessage{Id: "2", Name: "Geeking"}})

	// CRUD
	r.HandleFunc("/api/message", findAll).Methods("GET")         // findAll
	r.HandleFunc("/api/message/{id}", getMessage).Methods("GET") // get(id)
	r.HandleFunc("/api/message", create).Methods("POST")         // create
	r.HandleFunc("/api/message/{id}", update).Methods("PUT")     // update(id)
	r.HandleFunc("/api/message/{id}", delete).Methods("DELETE")  // delete(id)

	log.Fatal(http.ListenAndServe(":8080", r))
}
