package routes

import (
	"human-resources-backend/controllers"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RegisterUserRoutes(router fiber.Router, db *gorm.DB) {
	userController := controllers.NewUserController(db)

	user := router.Group("/users")

	user.Get("/", userController.GetUserList)
	user.Get(":id", userController.GetUserById)
	user.Post("/", userController.CreateUser)
	user.Put(":id", userController.UpdateUser)
	user.Delete(":id", userController.DeleteUser)
}
