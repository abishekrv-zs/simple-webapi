package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"simple-webapi/employee"
)

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/employee", employee.GetAllEmployee).Methods("GET")
	router.HandleFunc("/employee", employee.PostEmployee).Methods("POST")
	router.HandleFunc("/employee/{id}", employee.GetEmployee).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", router))
}
