package handler

import (
	"encoding/json"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"sandbox3.0/persistence/model"
)

func (wh *WebHandler) GetDepartments(c *fiber.Ctx) error {
	// get departments
	depts, err := wh.rs.GetDepartments()
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	// success
	response := make(map[string]interface{})
	response["success"] = true
	response["departments"] = depts

	// response
	jsonRes, _ := json.Marshal(response)
	c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)
	return c.Status(fiber.StatusOK).Send(jsonRes)
}

func (wh *WebHandler) GetDepartment(c *fiber.Ctx) error {
	id := c.Params("id")

	// request validation
	idInt, valErr := strconv.Atoi(id)
	if valErr != nil {
		return fiber.NewError(fiber.StatusBadRequest, valErr.Error())
	}

	// get department
	dept, err := wh.rs.GetDepartment(idInt)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	// success
	response := make(map[string]interface{})
	response["success"] = true
	response["department"] = dept

	// response
	jsonRes, _ := json.Marshal(response)
	c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)
	return c.Status(fiber.StatusOK).Send(jsonRes)
}

func (wh *WebHandler) CreateDepartment(c *fiber.Ctx) error {
	var dept model.Department

	// request validation
	valErr := c.BodyParser(&dept)
	if valErr != nil {
		return fiber.NewError(fiber.StatusBadRequest, valErr.Error())
	}

	// create department
	err := wh.rs.CreateDepartment(&dept)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	// success
	response := make(map[string]interface{})
	response["success"] = true
	response["department"] = dept

	// response
	jsonRes, _ := json.Marshal(response)
	c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)
	return c.Status(fiber.StatusOK).Send(jsonRes)
}

func (wh *WebHandler) UpdateDepartment(c *fiber.Ctx) error {
	id := c.Params("id")
	var dept model.Department

	// request validation
	idInt, val1Err := strconv.Atoi(id)
	if val1Err != nil {
		return fiber.NewError(fiber.StatusBadRequest, val1Err.Error())
	}
	val2Err := c.BodyParser(&dept)
	if val2Err != nil {
		return fiber.NewError(fiber.StatusBadRequest, val2Err.Error())
	}

	// update department
	dept.ID = idInt
	err := wh.rs.UpdateDepartment(&dept)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	// success
	response := make(map[string]interface{})
	response["success"] = true
	response["department"] = dept

	// response
	jsonRes, _ := json.Marshal(response)
	c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSON)
	return c.Status(fiber.StatusOK).Send(jsonRes)
}

func (wh *WebHandler) DeleteDepartment(c *fiber.Ctx) error {
	id := c.Params("id")
	var dept model.Department

	// request validation
	idInt, valErr := strconv.Atoi(id)
	if valErr != nil {
		return fiber.NewError(fiber.StatusBadRequest, valErr.Error())
	}

	// delete department
	dept.ID = idInt
	err := wh.rs.DeleteDepartment(&dept)
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
