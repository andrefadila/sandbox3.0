package task

import (
	"fmt"

	"sandbox3.0/persistence/model"
	"sandbox3.0/repository"
)

func Level3No3a(rs *repository.Service) {
	// create a new department
	department := &model.Department{
		Name: "IT A",
	}
	deptErr := rs.CreateDepartment(department)
	if deptErr != nil {
		fmt.Println("Error: ", deptErr.Error())
		return
	}

	// create 2 new employee
	employee1 := &model.Employee{
		Name:         "Fadila",
		DepartmentId: department.ID,
	}
	emp1Err := rs.CreateEmployee(employee1)
	if emp1Err != nil {
		fmt.Println("Error: ", emp1Err.Error())
	}

	employee2 := &model.Employee{
		Name:         "Andre",
		DepartmentId: department.ID,
	}
	emp2Err := rs.CreateEmployee(employee2)
	if emp2Err != nil {
		fmt.Println("Error: ", emp2Err.Error())
	}
}

func Level3No3b(rs *repository.Service) {
	// create a new department
	department := &model.Department{
		Name: "IT B",
	}
	deptErr := rs.CreateDepartment(department)
	if deptErr != nil {
		fmt.Println("Error: ", deptErr.Error())
		return
	}

	// get and update employee to new department
	emps, getEmpErr := rs.GetEmployees()
	if getEmpErr != nil {
		fmt.Println("Error: ", getEmpErr.Error())
		return
	}

	for _, emp := range emps {
		emp.DepartmentId = department.ID
		updateErr := rs.UpdateEmployee(&emp)
		if updateErr != nil {
			fmt.Println("Error: ", updateErr.Error())
		}
	}
}

func Level3No4(rs *repository.Service) {
	var depts []model.Department
	for i := 1; i <= 10; i++ {
		depts = append(depts, model.Department{Name: fmt.Sprintf("IT %d", i)})
	}

	// create department
	rs.CreateDepartments(depts)

	// create employees
	var emps []model.Employee
	for _, dept := range depts {
		for i := 1; i <= 10; i++ {
			emps = append(emps, model.Employee{Name: fmt.Sprintf("%s Employee %d ", dept.Name, i), DepartmentId: dept.ID})
		}
		rs.CreateEmployees(emps)
		emps = nil
	}
}
