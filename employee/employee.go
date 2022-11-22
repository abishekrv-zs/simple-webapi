package employee

import (
	"encoding/json"
	"io"
	"net/http"
)

type employee struct {
	Id      int    `json:"id"`
	Name    string `json:"name,omitempty"`
	Age     int    `json:"age,omitempty"`
	Address string `json:"address,omitempty"`
}

var employees = []employee{
	{1, "Abishek", 22, "Chennai"},
	{2, "Kavin", 22, "Mumbai"},
}

func GetEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	resp, _ := json.Marshal(employees)
	w.Write(resp)
}

func PostEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	var newEmp employee
	reqBody, _ := io.ReadAll(r.Body)

	if err := json.Unmarshal(reqBody, &newEmp); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "invalid request body"}`))
		return
	}

	if newEmp.Id == 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error": "id cannot be 0"}`))
		return
	}

	for _, emp := range employees {
		if newEmp.Id == emp.Id {
			w.WriteHeader(http.StatusConflict)
			w.Write([]byte(`{"error": "id already exists"}`))
			return
		}
	}

	employees = append(employees, newEmp)
	resp, _ := json.Marshal(newEmp)
	w.WriteHeader(http.StatusCreated)
	w.Write(resp)
}
