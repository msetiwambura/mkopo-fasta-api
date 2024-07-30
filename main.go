package main

import (
	"github.com/joho/godotenv"
	"loanapi/configs"
	"loanapi/models"
	"loanapi/routes"
	"log"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	configs.ConnectDB()
	models.MigrateDB()

	r := routes.SetupRouter() // Use the router returned by SetupRouter
	errr := r.Run(":8080")
	if errr != nil {
		log.Fatal("Error starting the server:", errr)
	}
}
