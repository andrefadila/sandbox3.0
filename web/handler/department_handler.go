package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"sandbox3.0/persistence/model"
)

func (wh *WebHandler) GetDepartments(w http.ResponseWriter, r *http.Request) {
	// get departments
	depts, err := wh.rs.GetDepartments()
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	// success
	response := make(map[string]interface{})
	response["success"] = true
	response["departments"] = depts

	// response
	jsonRes, _ := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonRes)
}

func (wh *WebHandler) GetDepartment(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	fmt.Println("GetDepartment id: ", id)

	// request validation
	idInt, valErr := strconv.Atoi(id)
	if valErr != nil {
		http.Error(w, valErr.Error(), http.StatusBadRequest)
		return
	}

	// get department
	dept, err := wh.rs.GetDepartment(idInt)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	// success
	response := make(map[string]interface{})
	response["success"] = true
	response["department"] = dept

	// response
	jsonRes, _ := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonRes)
}

func (wh *WebHandler) CreateDepartment(w http.ResponseWriter, r *http.Request) {
	var dept model.Department

	// request validation
	valErr := json.NewDecoder(r.Body).Decode(&dept)
	if valErr != nil {
		http.Error(w, valErr.Error(), http.StatusBadRequest)
		return
	}

	// create department
	err := wh.rs.CreateDepartment(&dept)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// success
	response := make(map[string]interface{})
	response["success"] = true
	response["department"] = dept

	// response
	jsonRes, _ := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonRes)
}

func (wh *WebHandler) UpdateDepartment(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	var dept model.Department

	// request validation
	idInt, val1Err := strconv.Atoi(id)
	if val1Err != nil {
		http.Error(w, val1Err.Error(), http.StatusBadRequest)
		return
	}
	val2Err := json.NewDecoder(r.Body).Decode(&dept)
	if val2Err != nil {
		http.Error(w, val2Err.Error(), http.StatusBadRequest)
		return
	}

	// update department
	dept.ID = idInt
	err := wh.rs.UpdateDepartment(&dept)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// success
	response := make(map[string]interface{})
	response["success"] = true
	response["department"] = dept

	// response
	jsonRes, _ := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonRes)
}

func (wh *WebHandler) DeleteDepartment(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	var dept model.Department

	// request validation
	idInt, valErr := strconv.Atoi(id)
	if valErr != nil {
		http.Error(w, valErr.Error(), http.StatusBadRequest)
		return
	}

	// create department
	dept.ID = idInt
	err := wh.rs.DeleteDepartment(&dept)
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
