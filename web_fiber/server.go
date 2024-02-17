package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"

	"sandbox3.0/persistence"
	"sandbox3.0/repository"
	"sandbox3.0/web_fiber/handler"
)

func main() {
	// check mysql connection
	db, dbErr := persistence.OpenMySqlConn()
	if dbErr != nil {
		fmt.Println("Database error: ", dbErr.Error())
	}
	defer db.Close()
	db.Automigrate()

	// initiate service
	rs := repository.NewService(db.MysqlDB)

	// initiate web handler
	app := fiber.New(fiber.Config{
		AppName: "Sandbox 3.0",
		ServerHeader: fiber.MIMEApplicationJSON,
    })
	wh := handler.NewWebHandler(rs, app)

	// start server
	wh.RegisterRoute()
	wh.App.Listen(":3030")
}
