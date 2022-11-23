package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/emp", employeeHandler)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
