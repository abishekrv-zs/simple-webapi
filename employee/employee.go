package employee

import (
	"encoding/json"
	"io"
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
	w.WriteHeader(http.StatusOK)

	resp, _ := json.Marshal(employees)

	if _, err := w.Write(resp); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(`{"error": "failed to write response"}`))
	}
}

func PostEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	var newEmp employee
	reqBody, _ := io.ReadAll(r.Body)

	if err := json.Unmarshal(reqBody, &newEmp); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(`{"error": "invalid request body"}`))
		return
	}

	employees = append(employees, newEmp)
	resp, _ := json.Marshal(newEmp)

	if _, err := w.Write(resp); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(`{"error": "failed to write response"}`))
	}
}
