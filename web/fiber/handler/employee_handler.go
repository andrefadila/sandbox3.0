package handler

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
	"sandbox3.0/persistence/model"
)

func (wh *WebHandler) GetEmployees(c *fiber.Ctx) error {
	// get employees
	emps, err := wh.rs.GetEmployees()
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	// response
	return c.JSON(fiber.Map{"success": true, "employees": emps})
}

func (wh *WebHandler) GetEmployee(c *fiber.Ctx) error {
	id := c.Params("id")

	// request validation
	idInt, valErr := strconv.Atoi(id)
	if valErr != nil {
		return fiber.NewError(fiber.StatusBadRequest, valErr.Error())
	}

	// get employee
	emp, err := wh.rs.GetEmployee(idInt)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	// response
	return c.JSON(fiber.Map{"success": true, "employee": emp})
}

func (wh *WebHandler) CreateEmployee(c *fiber.Ctx) error {
	var emp model.Employee

	// request validation
	valErr := c.BodyParser(&emp)
	if valErr != nil {
		return fiber.NewError(fiber.StatusBadRequest, valErr.Error())
	}

	// create employee
	err := wh.rs.CreateEmployee(&emp)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	// response
	return c.JSON(fiber.Map{"success": true, "employee": emp})
}

func (wh *WebHandler) UpdateEmployee(c *fiber.Ctx) error {
	var emp model.Employee

	// request validation
	id, val1Err := c.ParamsInt("id")
	if val1Err != nil {
		return fiber.NewError(fiber.StatusBadRequest, val1Err.Error())
	}
	val2Err := c.BodyParser(&emp)
	if val2Err != nil {
		return fiber.NewError(fiber.StatusBadRequest, val2Err.Error())
	}

	// update employee
	emp.ID = id
	err := wh.rs.UpdateEmployee(&emp)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	// response
	return c.JSON(fiber.Map{"success": true, "employee": emp})
}

func (wh *WebHandler) DeleteEmployee(c *fiber.Ctx) error {
	var emp model.Employee

	// request validation
	id, valErr := c.ParamsInt("id")
	if valErr != nil {
		return fiber.NewError(fiber.StatusBadRequest, valErr.Error())
	}

	// delete employee
	emp.ID = id
	err := wh.rs.DeleteEmployee(&emp)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	// response
	return c.JSON(fiber.Map{"success": true})
}
