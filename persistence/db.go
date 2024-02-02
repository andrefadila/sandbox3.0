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

func (s *DB) Close() error {
	db, err := s.mysqlDB.DB()
	if err != nil {
		return err
	}

	return db.Close()
}