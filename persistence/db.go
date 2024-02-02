package persistence

import "gorm.io/gorm"

type DB struct {
	mysqlDB *gorm.DB
}
