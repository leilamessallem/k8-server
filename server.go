package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

type Numbers struct {
	A float32 `json:"a"`
	B float32 `json:"b"`
}

func main() {
	router := mux.NewRouter()
	
	router.HandleFunc("/health", ping).Methods("GET")
	router.HandleFunc("/add", add).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", router))
}

func ping(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(time.Now().Format(time.RFC3339))
}

func add(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var numbers Numbers
	_ = json.NewDecoder(r.Body).Decode(&numbers)
	result := numbers.A + numbers.B
	json.NewEncoder(w).Encode(result)
}
