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
	mux.HandleFunc("/departments", wh.GetDepartments)
	mux.HandleFunc("/departments/{id}", wh.GetDepartment)
	mux.HandleFunc("POST /departments", wh.CreateDepartment)
	mux.HandleFunc("PUT /departments/{id}", wh.UpdateDepartment)
	mux.HandleFunc("DELETE /departments/{id}", wh.DeleteDepartment)
	return mux
}
