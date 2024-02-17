package handler

import (
	"encoding/json"
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

	// success
	response := make(map[string]interface{})
	response["success"] = true
	response["employees"] = emps

	// response
	jsonRes, _ := json.Marshal(response)
	c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)
	return c.Status(fiber.StatusOK).Send(jsonRes)
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

	// success
	response := make(map[string]interface{})
	response["success"] = true
	response["employee"] = emp

	// response
	jsonRes, _ := json.Marshal(response)
	c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)
	return c.Status(fiber.StatusOK).Send(jsonRes)
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

	// success
	response := make(map[string]interface{})
	response["success"] = true
	response["employee"] = emp

	// response
	jsonRes, _ := json.Marshal(response)
	c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)
	return c.Status(fiber.StatusOK).Send(jsonRes)
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

	// success
	response := make(map[string]interface{})
	response["success"] = true
	response["employee"] = emp

	// response
	jsonRes, _ := json.Marshal(response)
	c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)
	return c.Status(fiber.StatusOK).Send(jsonRes)
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

	// success
	response := make(map[string]interface{})
	response["success"] = true

	// response
	jsonRes, _ := json.Marshal(response)
	c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)
	return c.Status(fiber.StatusOK).Send(jsonRes)
}
