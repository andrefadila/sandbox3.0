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
	mux.HandleFunc("/department/{id}", wh.GetDepartment)
	return mux
}
