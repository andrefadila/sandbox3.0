package repository

import (
	"sandbox3.0/persistence/model"
)

func (s *Service) GetEmployee(id int) (*model.Employee, error) {
	err := s.db.Table("employees").Where("id = ?", id).First(&model.Employee{}).Error
	if err != nil {
		return nil, err
	}

	return &model.Employee{}, nil
}

func (s *Service) GetEmployees() ([]model.Employee, error) {
	var emps []model.Employee
	err := s.db.Table("employees").Limit(10).Scan(&emps).Error
	if err != nil {
		return nil, err
	}

	return emps, nil
}

func (s *Service) CreateEmployee(e *model.Employee) error {
	err := s.db.Create(&e).Error
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) CreateEmployees(emps []model.Employee) error {
	err := s.db.CreateInBatches(&emps, 5).Error
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) UpdateEmployee(e *model.Employee) error {
	err := s.db.Model(&e).Where("id = ?", e.ID).Updates(&e).Error
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) DeleteEmployee(e *model.Employee) error {
	err := s.db.Where("id = ?", e.ID).Delete(&e).Error
	if err != nil {
		return err
	}

	return nil
}
