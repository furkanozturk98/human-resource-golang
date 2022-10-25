package main

import (
	"fmt"
	"human-resources-backend/database"
	"human-resources-backend/routes"

	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Application is starting")

	godotenv.Load()

	db := database.ConnectDB()
	database.Migration(db)

	routes.RegisterRoutes(db)
}
