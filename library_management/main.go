package main

import (
	"library_management/controllers"
	"library_management/models"
	"library_management/services"
)

func main() {
	library := services.NewLibrary()

	library.Members[1] = &models.Member{ID: 1, Name: "Alice"}
	library.Members[2] = &models.Member{ID: 2, Name: "Bob"}

	controllers.StartConsole(library, library.Members)
}
