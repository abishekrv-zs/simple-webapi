package employee

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
	"strconv"
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

func GetAllEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	resp, err := json.Marshal(employees)
	if err != nil {
		log.Println(err)
		return
	}

	if _, err := w.Write(resp); err != nil {
		log.Println(err)
	}
}

func GetEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	reqEmpId, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		log.Println(err)
	}

	if reqEmpId <= 0 {
		w.WriteHeader(http.StatusBadRequest)
		if _, err := w.Write([]byte(`{"error": "invalid employee id"}`)); err != nil {
			log.Println(err)
		}
		return
	}

	for _, emp := range employees {
		if emp.Id == reqEmpId {
			respBody, err := json.Marshal(emp)
			if err != nil {
				log.Println(err)
			}
			if _, err := w.Write(respBody); err != nil {
				log.Println(err)
			}
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	if _, err := w.Write([]byte(`{"error": "employee id not found"}`)); err != nil {
		log.Println(err)
	}
}

func PostEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	var newEmp employee
	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
	}

	if err := json.Unmarshal(reqBody, &newEmp); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		if _, err := w.Write([]byte(`{"error": "invalid request body"}`)); err != nil {
			log.Println(err)
		}
		return
	}

	if newEmp.Id <= 0 {
		w.WriteHeader(http.StatusBadRequest)
		if _, err := w.Write([]byte(`{"error": "invalid id"}`)); err != nil {
			log.Println(err)
		}
		return
	}

	for _, emp := range employees {
		if newEmp.Id == emp.Id {
			w.WriteHeader(http.StatusConflict)
			if _, err := w.Write([]byte(`{"error": "id already exists"}`)); err != nil {
				log.Println(err)
			}
			return
		}
	}

	employees = append(employees, newEmp)
	resp, err := json.Marshal(newEmp)
	if err != nil {
		log.Println(err)
	}

	w.WriteHeader(http.StatusCreated)
	if _, err := w.Write(resp); err != nil {
		log.Println(err)
	}
}
