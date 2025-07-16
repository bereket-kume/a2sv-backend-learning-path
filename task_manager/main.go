package main

import (
	"task_manager/database"
	"task_manager/router"
)

func main() {
	database.ConnectDB()
	r := router.SetUpRoutes()
	r.Run()
}
