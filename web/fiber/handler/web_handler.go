package handler

import (
	"github.com/gofiber/fiber/v2"

	"sandbox3.0/repository"
)

type WebHandler struct {
	rs  *repository.Service
	App *fiber.App
}

func NewWebHandler(rs *repository.Service, app *fiber.App) *WebHandler {
	return &WebHandler{rs, app}
}

func (wh *WebHandler) Start() {
	// departments
	wh.App.Get("/departments", wh.GetDepartments)
	wh.App.Get("/departments/:id", wh.GetDepartment)
	wh.App.Post("/departments", wh.CreateDepartment)
	wh.App.Put("/departments/:id", wh.UpdateDepartment)
	wh.App.Delete("/departments/:id", wh.DeleteDepartment)
	// employees
	wh.App.Get("/employees", wh.GetEmployees)
	wh.App.Get("/employees/:id", wh.GetEmployee)
	wh.App.Post("/employees", wh.CreateEmployee)
	wh.App.Put("/employees/:id", wh.UpdateEmployee)
	wh.App.Delete("/employees/:id", wh.DeleteEmployee)

	// start listen
	wh.App.Listen(":3030")
}
