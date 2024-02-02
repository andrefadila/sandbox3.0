package employee

import (
	"gorm.io/gorm"
	"sandbox3.0/persistence/model"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db}
}

type EmployeeService interface {
	GetEmployee(id int) (*model.Employee, error)
	GetEmployees() ([]*model.Employee, error)
	CreateEmployee(e *model.Employee) error
	UpdateEmployee(e *model.Employee) error
	DeleteEmployee(id int) error
}

func (es *Repository) GetEmployee(id int) (*model.Employee, error) {
	err := es.db.Table("employees").Where("id = ?", id).First(&model.Employee{}).Error
	if err != nil {
		return nil, err
	}

	return &model.Employee{}, nil
}

func (es *Repository) CreateEmployee(e *model.Employee) error {
	err := es.db.Create(&e).Error
	if err != nil {
		return err
	}

	return nil
}

func (es *Repository) UpdateEmployee(e *model.Employee) error {
	err := es.db.Model(&e).Where("id = ?", e.ID).Updates(&e).Error
	if err != nil {
		return err
	}

	return nil
}

func (es *Repository) DeleteEmployee(e *model.Employee) error {
	err := es.db.Where("id = ?", e.ID).Delete(&e).Error
	if err != nil {
		return err
	}

	return nil
}
