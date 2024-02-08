package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"sandbox3.0/persistence"
	"sandbox3.0/pkg/department"
	"sandbox3.0/pkg/employee"
	"sandbox3.0/task"
)

func init() {
	// Load env file
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err.Error())
	}
}

func main() {
	// mysql db config
	dbHost := os.Getenv("MYSQL_HOST")
	dbPassword := os.Getenv("MYSQL_ROOT_PASSWORD")
	dbUser := os.Getenv("MYSQL_USER")
	dbName := os.Getenv("MYSQL_DATABASE")

	// check connection
	db, dbErr := persistence.OpenMySqlConn(dbUser, dbPassword, dbHost, dbName)
	if dbErr != nil {
		fmt.Println("Database error: ", dbErr.Error())
	}
	defer db.Close()
	db.Automigrate()

	// initiate service
	deptRepo := department.NewRepository(db.MysqlDB)
	empRepo := employee.NewRepository(db.MysqlDB)

	// task
	// task.Level3No3a(deptRepo, empRepo)
	// task.Level3No3b(deptRepo, empRepo)
	task.Level3No4(deptRepo, empRepo)
}
