package main

import (
	"fmt"

	"github.com/joho/godotenv"
	"sandbox3.0/persistence"
	"sandbox3.0/repository"
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
	// check connection
	db, dbErr := persistence.OpenMySqlConn()
	if dbErr != nil {
		fmt.Println("Database error: ", dbErr.Error())
	}
	defer db.Close()
	db.Automigrate()

	// initiate service
	rs := repository.NewService(db.MysqlDB)

	// task
	// task.Level3No3a(repo)
	// task.Level3No3b(repo)
	task.Level3No4(rs)
}
