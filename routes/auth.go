package routes

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"human-resources-backend/controllers"
)

func RegisterAuthRoutes(router fiber.Router, db *gorm.DB) {
	authController := controllers.NewAuthController(db)

	auth := router.Group("/auth")

	auth.Get("/token/new", authController.GetNewAccessToken)

}
