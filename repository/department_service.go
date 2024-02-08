package repository

import (
	"sandbox3.0/persistence/model"
)

func (s *Service) GetDepartment(id int) (*model.Department, error) {
	var dept model.Department
	err := s.db.Table("departments").Where("id = ?", id).First(&dept).Error
	if err != nil {
		return nil, err
	}

	return &dept, nil
}

func (s *Service) GetDepartments() ([]model.Department, error) {
	var dept []model.Department
	err := s.db.Table("departments").Limit(10).Scan(&dept).Error
	if err != nil {
		return nil, err
	}

	return dept, nil
}

func (s *Service) CreateDepartment(dept *model.Department) error {
	err := s.db.Create(&dept).Error
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) CreateDepartments(depts []model.Department) error {
	err := s.db.CreateInBatches(&depts, 5).Error
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) UpdateDepartment(dept *model.Department) error {
	err := s.db.Model(&dept).Where("id = ?", dept.ID).Updates(&dept).Error
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) DeleteDepartment(dept *model.Department) error {
	err := s.db.Where("id = ?", dept.ID).Delete(&dept).Error
	if err != nil {
		return err
	}

	return nil
}
