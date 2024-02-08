package main

import (
	"fmt"
	"net/http"

	"sandbox3.0/persistence"
	"sandbox3.0/repository"
	"sandbox3.0/web/handler"
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
	wh := handler.NewWebHandler(rs)

	// start server
	http.ListenAndServe(":3030", wh.RouteHandler())
}
