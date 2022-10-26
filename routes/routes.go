package routes

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"human-resources-backend/configs"
	"log"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RegisterRoutes(db *gorm.DB, awsSession *session.Session) {
	app := fiber.New()
	//api := app.Group(configs.API_VERSION, middleware.SetContentTypeJSON)
	api := app.Group(configs.API_VERSION)

	RegisterAuthRoutes(api, db)
	RegisterUserRoutes(api, db)
	RegisterEmployeeRoutes(api, db)
	RegisterCompanyRoutes(api, db, awsSession)

	if configs.PORT == "" {
		configs.PORT = "8080"
	}

	log.Fatalf("Server error: %v", app.Listen(":8080"))
}
