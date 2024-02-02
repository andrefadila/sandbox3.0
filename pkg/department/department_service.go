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
	CreateDepartments(d []*model.Department) error
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

func (es *Repository) GetDepartments() ([]model.Department, error) {
	var dept []model.Department
	err := es.db.Table("departments").Limit(10).Scan(&dept).Error
	if err != nil {
		return nil, err
	}

	return dept, nil
}

func (r *Repository) CreateDepartment(dept *model.Department) error {
	err := r.db.Create(&dept).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) CreateDepartments(depts []*model.Department) error {
	err := r.db.CreateInBatches(&depts, 10).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) UpdateDepartment(dept *model.Department) error {
	err := r.db.Model(&dept).Where("id = ?", dept.ID).Updates(&dept).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) DeleteDepartment(dept *model.Department) error {
	err := r.db.Where("id = ?", dept.ID).Delete(&dept).Error
	if err != nil {
		return err
	}

	return nil
}
