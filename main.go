package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

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
	json.NewEncoder(w).Encode(message)

}

func GetHandler(w http.ResponseWriter, r *http.Request) {
	var messages []Message
	DB.Find(&messages)
	json.NewEncoder(w).Encode(messages)
	for _, message := range messages {
		fmt.Fprintf(
			w, "ID: %d, CreatedAt: %s, Task: %s, Is Done: %t\n",
			message.ID, message.CreatedAt, message.Task, message.IsDone)
	}
}

func main() {
	InitDB()
	DB.AutoMigrate(&Message{})
	router := mux.NewRouter()
	router.HandleFunc("/api/all_tasks", GetHandler).Methods("GET")
	router.HandleFunc("/api/new_task", PostHandler).Methods("POST")
	http.ListenAndServe("localhost:8080", router)
}
