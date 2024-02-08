package persistence

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func OpenMySqlConn(dbUser, dbPassword, dbHost, dbName string) (*DB, error) {
	dsn := dbUser + ":" + dbPassword + "@tcp(" + dbHost + ":3306)/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{CreateBatchSize: 5})
	if err != nil {
		return nil, err
	}

	db.Session(&gorm.Session{CreateBatchSize: 5})

	return &DB{
		MysqlDB: db,
	}, nil
}
