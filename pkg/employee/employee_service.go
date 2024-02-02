package employee

import (
	"gorm.io/gorm"
	"sandbox3.0/persistence/model"
)

type EmployeeRepository struct {
	db *gorm.DB
}

func NewEmployeeRepository(db *gorm.DB) *EmployeeRepository {
	return &EmployeeRepository{db}
}

type EmployeeService interface {
	GetEmployee(id int) (*model.Employee, error)
	GetEmployees() ([]*model.Employee, error)
	CreateEmployee(e *model.Employee) error
	UpdateEmployee(e *model.Employee) error
	DeleteEmployee(id int) error
}

func (es *EmployeeRepository) GetEmployee(id int) (*model.Employee, error) {
	err := es.db.Table("employees").Where("id = ?", id).First(&model.Employee{}).Error
	if err != nil {
		return nil, err
	}

	return &model.Employee{}, nil
}

func (es *EmployeeRepository) CreateEmployee(e *model.Employee) error {
	err := es.db.Create(&e).Error
	if err != nil {
		return err
	}

	return nil
}
