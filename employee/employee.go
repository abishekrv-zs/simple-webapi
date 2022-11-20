package employee

import (
	"encoding/json"
	"net/http"
)

type employee struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Address string `json:"address"`
}

var employees = []employee{
	{1, "Abishek", 22, "Chennai"},
	{2, "Kavin", 22, "Mumbai"},
}

func GetEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(employees); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func PostEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	var newEmp employee
	if err := json.NewDecoder(r.Body).Decode(&newEmp); err != nil {
		http.Error(w, "Conflict in parsing request body", http.StatusConflict)
		return
	}

	w.WriteHeader(http.StatusCreated)
	employees = append(employees, newEmp)

	if err := json.NewEncoder(w).Encode(newEmp); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
