package main

import (
	"main/database"
	router "main/routers"
	"main/service"
)

func main() {
	// init db
	database.ConnectDB()

	// init service
	service.NewService()

	// init routes
	router.InitRoute()
}
