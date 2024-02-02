package task

import (
	"fmt"

	"sandbox3.0/persistence/model"
	"sandbox3.0/pkg/department"
)

func Level3No1(deptRepo *department.Repository) {
	// create a new department
	department := &model.Department{
		Name: "IT",
	}
	err := deptRepo.CreateDepartment(department)
	if err != nil {
		fmt.Println("Error: ", err.Error())
	}
}
