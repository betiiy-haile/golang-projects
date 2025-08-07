package main

import (
	"library-management/controllers"
	"library-management/models"
	"library-management/services"
)

func main() {
	library := services.NewLibrary()

	// Seed with a few members
	library.Members[1] = &models.Member{Id: 1, Name: "Alice"}
	library.Members[2] = &models.Member{Id: 2, Name: "Bob"}

	controllers.StartConsole(library)
}