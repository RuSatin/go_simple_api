package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type task struct {
	Message string
}

func GetHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, Go!")
}

func PostHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, task{Message: "Hello task"})
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/api/hello", GetHandler).Methods("GET")
	router.HandleFunc("/api/task", PostHandler).Methods("POST")
	http.ListenAndServe("localhost:8080", router)
}
