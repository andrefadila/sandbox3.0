package department

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

type Service interface {
	GetDepartment(id int) (*model.Department, error)
	GetDepartments() ([]*model.Department, error)
	CreateDepartment(d *model.Department) error
	UpdateDepartment(d *model.Department) error
	DeleteDepartment(id int) error
}

func (r *Repository) GetDepartment(id int) (*model.Department, error) {
	err := r.db.Table("departments").Where("id = ?", id).First(&model.Department{}).Error
	if err != nil {
		return nil, err
	}

	return &model.Department{}, nil
}

func (r *Repository) CreateDepartment(e *model.Department) error {
	err := r.db.Create(&e).Error
	if err != nil {
		return err
	}

	return nil
}
