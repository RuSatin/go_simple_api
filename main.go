package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

var task string

type requestBody struct {
	ID      uint   `gorm:"primarykey"`
	Message string `json:"message"`
}

func PostHandler(w http.ResponseWriter, r *http.Request) {
	var request requestBody
	json.NewDecoder(r.Body).Decode(&request)

	message := Message{
		Task:   request.Message,
		IsDone: true, // по умолчанию задача не выполнена
	}

	DB.Create(&message)

}

func GetHandler(w http.ResponseWriter, r *http.Request) {
	var messages []Message
	DB.Find(&messages)
	for _, message := range messages {
		fmt.Fprintf(
			w, "ID: %d, CreatedAt: %s, Task: %s, Is Done: %t\n",
			message.ID, message.CreatedAt, message.Task, message.IsDone)
	}
}

func main() {
	InitDB()
	DB.Create(task)
	DB.AutoMigrate(&Message{})

	router := mux.NewRouter()
	router.HandleFunc("/api/hello", GetHandler).Methods("GET")
	router.HandleFunc("/api/task", PostHandler).Methods("POST")
	http.ListenAndServe("localhost:8080", router)
}
