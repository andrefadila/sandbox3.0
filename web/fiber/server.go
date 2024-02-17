package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"

	"sandbox3.0/persistence"
	"sandbox3.0/repository"
	"sandbox3.0/web/fiber/handler"
)

func main() {
	// check mysql connection
	db, dbErr := persistence.OpenMySqlConn()
	if dbErr != nil {
		fmt.Println("Database error: ", dbErr.Error())
	}
	defer db.Close()
	db.Automigrate()

	// initiate db service
	rs := repository.NewService(db.MysqlDB)

	// initiate app & start
	app := fiber.New(fiber.Config{
		AppName: "Sandbox 3.0",
    })
	wh := handler.NewWebHandler(rs, app)
	wh.Start()
}
