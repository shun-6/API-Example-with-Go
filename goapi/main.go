package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Person struct {
	Name    string `json:"name"`
    Age     int    `json:"age"`
    Country string `json:"country"` 

}

func getPersonInfo(w http.ResponseWriter, r *http.Request) {
	person := Person{Name:"Emir", Age: 21, Country:"Türkiye"}

	response, err := json.Marshal(person)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)

}

func main(){

	http.HandleFunc("/person", getPersonInfo)

	log.Println("Server listening on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

}