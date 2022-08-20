package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func HandelRouting() {
	r := mux.NewRouter()
	r.HandleFunc("/employee", CreateEmployee).Methods("POST")
	r.HandleFunc("/employees", GetEmployees).Methods("GET")
	r.HandleFunc("/employeebyid/{eid}", GetEmployeeById).Methods("POST")
	r.HandleFunc("/employeebyid/{eid}", UpdateEmployee).Methods("PUT")
	r.HandleFunc("/employeebyid/{eid}", DeleteEmployeeById).Methods("DELETE")

	//Open API Port
	log.Fatal(http.ListenAndServe(":8080", r))

}
