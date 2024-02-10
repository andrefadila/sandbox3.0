package handler

import (
	"net/http"

	"sandbox3.0/repository"
)

type WebHandler struct {
	rs *repository.Service
}

func NewWebHandler(rs *repository.Service) *WebHandler {
	return &WebHandler{rs}
}

func (wh *WebHandler) RouteHandler() http.Handler {
	mux := http.NewServeMux()
	// departments
	mux.HandleFunc("/departments", wh.GetDepartments)
	mux.HandleFunc("/departments/{id}", wh.GetDepartment)
	mux.HandleFunc("POST /departments", wh.CreateDepartment)
	mux.HandleFunc("PUT /departments/{id}", wh.UpdateDepartment)
	mux.HandleFunc("DELETE /departments/{id}", wh.DeleteDepartment)
	// employees
	mux.HandleFunc("/employees", wh.GetEmployees)
	mux.HandleFunc("/employees/{id}", wh.GetEmployee)
	mux.HandleFunc("POST /employees", wh.CreateEmployee)
	mux.HandleFunc("PUT /employees/{id}", wh.UpdateEmployee)
	mux.HandleFunc("DELETE /employees/{id}", wh.DeleteEmployee)
	return mux
}
