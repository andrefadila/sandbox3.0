package department

import (
	"gorm.io/gorm"
	"sandbox3.0/persistence/model"
)

type DepartmentRepository struct {
	db *gorm.DB
}

func NewDepartmentRepository(db *gorm.DB) *DepartmentRepository {
	return &DepartmentRepository{db}
}

type DepartmentService interface {
	GetDepartment(id int) (*model.Department, error)
	GetDepartments() ([]*model.Department, error)
	CreateDepartment(d *model.Department) error
	UpdateDepartment(d *model.Department) error
	DeleteDepartment(id int) error
}

func (es *DepartmentRepository) GetDepartment(id int) (*model.Department, error) {
	err := es.db.Table("department").Where("id = ?", id).First(&model.Department{}).Error
	if err != nil {
		return nil, err
	}

	return &model.Department{}, nil
}

func (es *DepartmentRepository) CreateDepartment(e *model.Department) error {
	err := es.db.Create(&e).Error
	if err != nil {
		return err
	}

	return nil
}
