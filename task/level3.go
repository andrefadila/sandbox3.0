package task

import (
	"fmt"

	"sandbox3.0/persistence/model"
	"sandbox3.0/pkg/department"
	"sandbox3.0/pkg/employee"
)

func Level3No1(deptRepo *department.Repository, empRepo *employee.Repository) {
	// create a new department
	department := &model.Department{
		Name: "IT",
	}
	deptErr := deptRepo.CreateDepartment(department)
	if deptErr != nil {
		fmt.Println("Error: ", deptErr.Error())
		return
	}

	// create new employee
	employee1 := &model.Employee{
		Name: "Fadila",
		DepartmentId: department.ID,
	}
	emp1Err := empRepo.CreateEmployee(employee1)
	if emp1Err != nil {
		fmt.Println("Error: ", emp1Err.Error())
	}

	employee2 := &model.Employee{
		Name: "Andre",
		DepartmentId: department.ID,
	}
	emp2Err := empRepo.CreateEmployee(employee2)
	if emp2Err != nil {
		fmt.Println("Error: ", emp2Err.Error())
	}
}
