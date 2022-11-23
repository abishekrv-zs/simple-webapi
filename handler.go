package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func employeeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	switch r.Method {

	case "GET":
		empId := r.URL.Query().Get("id")

		// `/emp`
		if empId == "" {
			respBody, err := json.Marshal(employees)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			if _, err := w.Write(respBody); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			return
		}

		// `/emp?id=2`
		for _, emp := range employees {
			if emp.Id == empId {
				respBody, err := json.Marshal(emp)
				if err != nil {
					log.Println(err)
					w.WriteHeader(http.StatusInternalServerError)
					return
				}
				if _, err := w.Write(respBody); err != nil {
					log.Println(err)
					w.WriteHeader(http.StatusInternalServerError)
					return
				}
				return
			}
		}

		// `/emp?id=<invalid_id>`
		w.WriteHeader(http.StatusNotFound)
		if _, err := w.Write([]byte(`{"error":"id not found"}`)); err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

	case "POST":
		reqBody, err := io.ReadAll(r.Body)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
		}

		var newEmp employee
		if err := json.Unmarshal(reqBody, &newEmp); err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			if _, err := w.Write([]byte(`{"error":"invalid request body"}`)); err != nil {
				log.Println(err)
				return
			}
			return
		}

		employees = append(employees, newEmp)

		respBody, err := json.Marshal(newEmp)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		if _, err := w.Write(respBody); err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		if _, err := w.Write([]byte(`{"error":"invalid request method"}`)); err != nil {
			log.Println(err)
		}
	}
}
