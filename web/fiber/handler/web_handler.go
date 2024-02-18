package handler

import (
	jwtware "github.com/gofiber/contrib/jwt"
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
	// Login route
	wh.App.Post("/login", wh.Login)

	// auth middleware
	// basicAuthMiddleware(wh.App)
	jwtMiddleware(wh.App)

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

/**
 * Commented out because it's not used in this example
 */
// func basicAuthMiddleware(app *fiber.App) {
// 	app.Use(basicauth.New(basicauth.Config{
// 		Realm: "Forbidden",
// 		Authorizer: func(user, pass string) bool {
// 			if user == "admin" && pass == "12345" {
// 				return true
// 			}
// 			return false
// 		},
// 		Unauthorized: func(c *fiber.Ctx) error {
// 			return c.SendStatus(fiber.StatusUnauthorized)
// 		},
// 		ContextUsername: "_user",
// 		ContextPassword: "_pass",
// 	}))
// }

func jwtMiddleware(app *fiber.App) {
	app.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte("secret")},
	}))
}
