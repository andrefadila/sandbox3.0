package model

type Employee struct {
	ID           int    `gorm:"primary_key;auto_increment" json:"id"`
	Name         string `gorm:"size:128;not null;" json:"name"`
	DepartmentId int    `gorm:"null;" json:"departmentId"`
}
