package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func CreateTask(w http.ResponseWriter, r *http.Request) {
	var task Task
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result := DB.Create(&task)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}

func GetTask(w http.ResponseWriter, r *http.Request) {
	var tasks []Task
	result := DB.Find(&tasks)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	taskID := vars["id"]

	var task Task
	result := DB.First(&task, taskID)
	if result.Error != nil {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}

	var updateData map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&updateData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if taskText, ok := updateData["task"].(string); ok {
		task.Task = taskText
	}
	if isDone, ok := updateData["is_done"].(bool); ok {
		task.IsDone = isDone
	}

	DB.Save(&task)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}

func DeleteTask(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	taskID := vars["id"]

	var task Task
	result := DB.First(&task, taskID)
	if result.Error != nil {
		http.Error(w, "Task not found", http.StatusNotFound)
		return
	}

	DB.Delete(&task)

	w.WriteHeader(http.StatusNoContent)
}

func main() {
	InitDB()

	DB.AutoMigrate(&Task{})

	router := mux.NewRouter()
	router.HandleFunc("/api/tasks", CreateTask).Methods("POST")
	router.HandleFunc("/api/tasks", GetTask).Methods("GET")
	router.HandleFunc("/api/tasks/{id}", UpdateTask).Methods("PATCH")  // Обновление задачи
	router.HandleFunc("/api/tasks/{id}", DeleteTask).Methods("DELETE") //
	http.ListenAndServe(":8080", router)
}
