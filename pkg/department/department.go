package department

import (
	"sandbox3.0/persistence/model"
)

type DepartmentService interface {
	GetDepartment(id int) (*model.Department, error)
	GetDepartments() ([]*model.Department, error)
	CreateDepartment(d *model.Department) error
	UpdateDepartment(d *model.Department) error
	DeleteDepartment(id int) error
}
