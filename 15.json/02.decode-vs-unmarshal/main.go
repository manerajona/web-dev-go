package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type city struct {
	Bali       string  `json:"Postal"`
	Kauai      float64 `json:"Latitude"`
	Maui       float64 `json:"Longitude"`
	Java       string  `json:"Address"`
	NewZealand string  `json:"City"`
	Skye       string  `json:"State"`
	Oahu       string  `json:"Zip"`
	Hawaii     string  `json:"Country"`
}

type cities []city

func unmarshal(payload string, data *cities) {
	if err := json.Unmarshal([]byte(payload), &data); err != nil {
		log.Fatalln(err)
	}
}

func decode(r *http.Request, data *cities) {
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		log.Fatalln(err)
	}
}

func main() {
	var data cities

	payload := `[{"Postal":"zip","Latitude":37.7668,"Longitude":-122.3959,"Address":"","City":"SAN FRANCISCO","State":"CA","Zip":"94107","Country":"US"},{"Postal":"zip","Latitude":37.371991,"Longitude":-122.02602,"Address":"","City":"SUNNYVALE","State":"CA","Zip":"94085","Country":"US"}]`

	unmarshal(payload, &data)
	fmt.Println(data)
	fmt.Println(data[1].Kauai) // Latitude

	request, _ := http.NewRequest("GET", "/foo", bytes.NewBuffer([]byte(payload)))
	decode(request, &data)
	fmt.Println(data)
	fmt.Println(data[1].Kauai) // Latitude
}
