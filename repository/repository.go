package repository

import (
	"gorm.io/gorm"
	"sandbox3.0/persistence/model"
)

type Service struct {
	db *gorm.DB
}

func NewService(db *gorm.DB) *Service {
	return &Service{db}
}

type DepartmentService interface {
	GetDepartment(id int) (*model.Department, error)
	GetDepartments() ([]*model.Department, error)
	CreateDepartment(d *model.Department) error
	CreateDepartments(d []model.Department) error
	UpdateDepartment(d *model.Department) error
	DeleteDepartment(id int) error
}

type EmployeeService interface {
	GetEmployee(id int) (*model.Employee, error)
	GetEmployees() ([]model.Employee, error)
	CreateEmployee(e *model.Employee) error
	CreateEmployees(e []model.Employee) error
	UpdateEmployee(e *model.Employee) error
	DeleteEmployee(id int) error
}
