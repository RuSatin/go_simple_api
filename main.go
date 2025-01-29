package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

var task string

type requestBody struct {
	Message string `json:"message"`
}

func PostHandler(w http.ResponseWriter, r *http.Request) {
	var request requestBody

	json.NewDecoder(r.Body).Decode(&request)

	task = request.Message
	fmt.Fprintf(w, "Task %s", task)
}

func GetHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, Go!\n")
	fmt.Fprintf(w, "Your task is %s!", task)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/api/hello", GetHandler).Methods("GET")
	router.HandleFunc("/api/task", PostHandler).Methods("POST")
	http.ListenAndServe("localhost:8080", router)
}
