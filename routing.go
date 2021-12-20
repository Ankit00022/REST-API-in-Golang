package main

import (
	"log"
	"net/http"
	"pkg/mod/github.com/labstack/echo@v3.3.10+incompatible"
)

func HandlerRouting() {
	r := echo.NewRouter()
	r.Handlefunc("/employees", GetEmployees).Methods("GET")
	r.Handlefunc("/employee/{ePhoneNo}", GetEmployeeByPhoneNo).Methods("GET")
	r.Handlefunc("/employee", CreateEmployee).Methods("POST")
	r.Handlefunc("/employee/{ePhoneNo}", UpdateEmployee).Methods("PUT")
	r.Handlefunc("/employee/{ePhoneNo}", DeleteEmployee).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8080", r))

}
