package routes

import (
	"human-resources-backend/configs"
	"log"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RegisterRoutes(db *gorm.DB) {
	app := fiber.New()
	//api := app.Group(configs.API_VERSION, middleware.SetContentTypeJSON)
	api := app.Group(configs.API_VERSION)

	RegisterAuthRoutes(api, db)
	RegisterUserRoutes(api, db)
	RegisterEmployeeRoutes(api, db)
	RegisterCompanyRoutes(api, db)

	if configs.PORT == "" {
		configs.PORT = "8080"
	}

	log.Fatalf("Server error: %v", app.Listen(":8080"))
}
