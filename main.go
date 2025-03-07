package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

type RequestBody struct {
	Task string `json:"Hello"`
}

func GetHandler(w http.ResponseWriter, r *http.Request) {
	var tasks []Task
	result := DB.Find(&tasks)
	if result.Error != nil {
		http.Error(w, result.Error.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

func PostHandler(w http.ResponseWriter, r *http.Request) {
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

func main() {
	InitDB()

	DB.AutoMigrate(&Task{})

	router := mux.NewRouter()
	router.HandleFunc("/task/get", GetHandler).Methods("GET")
	router.HandleFunc("/task/post", PostHandler).Methods("POST")
	http.ListenAndServe(":8080", router)
}
