package employee

import (
	"sandbox3.0/persistence/model"
)

type EmployeeService interface {
	GetEmployee(id int) (*model.Employee, error)
	GetEmployees() ([]*model.Employee, error)
	CreateEmployee(e *model.Employee) error
	UpdateEmployee(e *model.Employee) error
	DeleteEmployee(id int) error
}
