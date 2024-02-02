package task

import (
	"fmt"

	"sandbox3.0/persistence/model"
	"sandbox3.0/pkg/department"
	"sandbox3.0/pkg/employee"
)

// type Level3Data struct {
// 	department model.Department
// 	employees  []model.Employee
// }

func Level3No3a(deptRepo *department.Repository, empRepo *employee.Repository) {
	// create a new department
	department := &model.Department{
		Name: "IT A",
	}
	deptErr := deptRepo.CreateDepartment(department)
	if deptErr != nil {
		fmt.Println("Error: ", deptErr.Error())
		return
	}

	// create 2 new employee
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

func Level3No3b(deptRepo *department.Repository, empRepo *employee.Repository) {
	// create a new department
	department := &model.Department{
		Name: "IT B",
	}
	deptErr := deptRepo.CreateDepartment(department)
	if deptErr != nil {
		fmt.Println("Error: ", deptErr.Error())
		return
	}

	// get and update employee to new department
	emps, getEmpErr := empRepo.GetEmployees()
	if getEmpErr != nil {
		fmt.Println("Error: ", getEmpErr.Error())
		return
	}
	
	for _, emp := range emps {
		emp.DepartmentId = department.ID
		updateErr := empRepo.UpdateEmployee(&emp)
		if updateErr != nil {
			fmt.Println("Error: ", updateErr.Error())
		}
	}
}
