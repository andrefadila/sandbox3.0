package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"sandbox3.0/persistence/model"
)

func (wh *WebHandler) GetEmployees(w http.ResponseWriter, r *http.Request) {
	// get employees
	emps, err := wh.rs.GetEmployees()
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	// success
	response := make(map[string]interface{})
	response["success"] = true
	response["employees"] = emps

	// response
	jsonRes, _ := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonRes)
}

func (wh *WebHandler) GetEmployee(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	// request validation
	idInt, valErr := strconv.Atoi(id)
	if valErr != nil {
		http.Error(w, valErr.Error(), http.StatusBadRequest)
		return
	}

	// get employee
	emp, err := wh.rs.GetEmployee(idInt)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	// success
	response := make(map[string]interface{})
	response["success"] = true
	response["employee"] = emp

	// response
	jsonRes, _ := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonRes)
}

func (wh *WebHandler) CreateEmployee(w http.ResponseWriter, r *http.Request) {
	var emp model.Employee

	// request validation
	valErr := json.NewDecoder(r.Body).Decode(&emp)
	if valErr != nil {
		http.Error(w, valErr.Error(), http.StatusBadRequest)
		return
	}

	// create employee
	err := wh.rs.CreateEmployee(&emp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// success
	response := make(map[string]interface{})
	response["success"] = true
	response["employee"] = emp

	// response
	jsonRes, _ := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonRes)
}

func (wh *WebHandler) UpdateEmployee(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	var emp model.Employee

	// request validation
	idInt, val1Err := strconv.Atoi(id)
	if val1Err != nil {
		http.Error(w, val1Err.Error(), http.StatusBadRequest)
		return
	}
	val2Err := json.NewDecoder(r.Body).Decode(&emp)
	if val2Err != nil {
		http.Error(w, val2Err.Error(), http.StatusBadRequest)
		return
	}

	// update employee
	emp.ID = idInt
	err := wh.rs.UpdateEmployee(&emp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// success
	response := make(map[string]interface{})
	response["success"] = true
	response["employee"] = emp

	// response
	jsonRes, _ := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonRes)
}

func (wh *WebHandler) DeleteEmployee(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	var emp model.Employee

	// request validation
	idInt, valErr := strconv.Atoi(id)
	if valErr != nil {
		http.Error(w, valErr.Error(), http.StatusBadRequest)
		return
	}

	// delete employee
	emp.ID = idInt
	err := wh.rs.DeleteEmployee(&emp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// success
	response := make(map[string]interface{})
	response["success"] = true

	// response
	jsonRes, _ := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonRes)
}
