package persistence

import (
	"errors"
	"fmt"

	"gorm.io/gorm"
	"sandbox3.0/persistence/model"
)

type DB struct {
	MysqlDB *gorm.DB
}

func (db *DB) Automigrate() error {
	return db.MysqlDB.AutoMigrate(
		&model.Employee{},
		&model.Department{},
	)
}

func (db *DB) MigrateAndSeed() error {
	// Migrate the schema. This is an example for departments.
	var departments model.Department
	if db.MysqlDB.Migrator().HasTable(departments) {
		db.MysqlDB.Migrator().DropTable(departments)
	}
	var employees model.Employee
	if db.MysqlDB.Migrator().HasTable(employees) {
		db.MysqlDB.Migrator().DropTable(employees)
	}
	if err := db.MysqlDB.AutoMigrate(departments, employees); err != nil {
		return fmt.Errorf("failed to migrate departments and employees: %w", err)
	}

	// Validate the table creation.
	if !db.MysqlDB.Migrator().HasTable(departments) {
		return errors.New("failed to create departments table")
	}
	if !db.MysqlDB.Migrator().HasTable(employees) {
		return errors.New("failed to create employees table")
	}

	// When we have the table, we can create some records (seed data).
	if db.MysqlDB.Migrator().HasTable(departments) {
		if err := db.MysqlDB.Model(departments).CreateInBatches([]model.Department{
			{
				Name: "Engineering",
			},
			{
				Name: "IT",
			},
			{
				Name: "Customer Service",
			},
		}, 10).Error; err != nil {
			return fmt.Errorf("failed to seed departments: %w", err)
		}
	}
	// When we have the table, we can create some records (seed data).
	if db.MysqlDB.Migrator().HasTable(employees) {
		if err := db.MysqlDB.Model(employees).CreateInBatches([]model.Employee{
			{
				Name: "Fadila",
			},
			{
				Name: "Andre",
			},
			{
				Name: "Mulyanto",
			},
		}, 10).Error; err != nil {
			return fmt.Errorf("failed to seed employees: %w", err)
		}
	}
	return nil
}

func (s *DB) Close() error {
	db, err := s.MysqlDB.DB()
	if err != nil {
		return err
	}

	return db.Close()
}
