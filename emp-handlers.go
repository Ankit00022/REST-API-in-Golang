package main

import (
	"encoding/json"
	"net/http"
)

func CreateEmployee(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("content-Type", "application/json")
	var emp Employee
	json.NewDecoder(r.Body).Decode(&emp)
	Database.Create(&emp)
	json.NewEncoder(w).Encode(emp)

}

func GetEmployees(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-Type", "application/json")
	var employees []Employee
	Database.Find(&employees)
	json.NewEncoder(w).Encode(employees)

}

func GetEmployeeByPhoneNo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-Type", "application/json")
	var employee Employee
	Database.First(&employee, echo.vars(r)["ePhoneNo"])
	json.NewEncoder(w).Encode(employee)
}

func UpdateEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-Type", "application/json")
	var employee Employee
	Database.First(&employee, echo.vars(r)["ePhoneNo"])
	json.NewDecoder(r.Body).Decode(&employee)
	Database.Save(&employee)
	json.NewEncoder(w).Encode(employee)

}

func DeleteEmployee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-Type", "application/json")
	var employee Employee
	Database.Delete(&employee, echo.vars(r)["ePhoneNo"])
	json.NewEncoder(w).Encode("employee is deleted !!")

}
