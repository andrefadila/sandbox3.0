package persistence

import (
	"gorm.io/gorm"
	"sandbox3.0/persistence/model"
)

type DB struct {
	mysqlDB *gorm.DB
}

func (s *DB) Automigrate() error {
	return s.mysqlDB.AutoMigrate(
		&model.Employee{},
		&model.Department{},
	)
}
