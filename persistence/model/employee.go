package model

type Employee struct {
	ID           int    `gorm:"primary_key;auto_increment"`
	Name         string `gorm:"size:128;not null;"`
	DepartmentId int    `gorm:"null;"`
}
