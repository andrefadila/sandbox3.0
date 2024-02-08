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
	mux.HandleFunc("/department", wh.GetDepartments)
	mux.HandleFunc("/department/{id}", wh.GetDepartment)
	mux.HandleFunc("POST /department", wh.CreateDepartment)
	mux.HandleFunc("PUT /department/{id}", wh.UpdateDepartment)
	mux.HandleFunc("DELETE /department/{id}", wh.DeleteDepartment)
	return mux
}
