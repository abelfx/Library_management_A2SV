package main

import (
	"library_management/controllers"
	"library_management/models"
	"library_management/services"
)

func main() {
	lib := services.NewLibrary()

	lib.Members[1] = models.Member{ID: 1, Name: "Natnael"}
	lib.Members[2] = models.Member{ID: 2, Name: "Betre"}

	controllers.StartConsole(lib)
}
