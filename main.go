package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"simple-webapi/employee"
)

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/employee", employee.GetEmployee).Methods("GET")
	router.HandleFunc("/employee", employee.PostEmployee).Methods("POST")

	log.Fatal(http.ListenAndServe(":8000", router))
}
