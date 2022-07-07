package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type person struct {
	Fname string   `json:"first_name,omitempty"`
	Lname string   `json:"last_name,omitempty"`
	Items []string `json:"items"`
}

func marshal(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	p1 := person{
		Fname: "James",
		Lname: "Bond",
		Items: []string{"Suit", "Gun", "Wry sense of humor"},
	}
	j, err := json.Marshal(p1)
	if err != nil {
		log.Println(err)
	}
	w.Write(j)
}

func encode(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	p1 := person{
		Fname: "James",
		Lname: "Bond",
		Items: []string{"Suit", "Gun", "Wry sense of humor"},
	}
	err := json.NewEncoder(w).Encode(p1)
	if err != nil {
		log.Println(err)
	}
}

func main() {
	http.HandleFunc("/marshal", marshal)
	http.HandleFunc("/encode", encode)
	http.ListenAndServe(":8080", nil)
}
