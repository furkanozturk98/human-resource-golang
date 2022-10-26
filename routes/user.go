package routes

import (
	"human-resources-backend/controllers"
	"human-resources-backend/middlewares"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RegisterUserRoutes(router fiber.Router, db *gorm.DB) {
	userController := controllers.NewUserController(db)

	user := router.Group("/users")

	user.Get("/", middlewares.JWTProtected(), userController.GetUserList)
	user.Get(":id", middlewares.JWTProtected(), userController.GetUserById)
	user.Post("/", middlewares.JWTProtected(), userController.CreateUser)
	user.Put(":id", middlewares.JWTProtected(), userController.UpdateUser)
	user.Delete(":id", middlewares.JWTProtected(), userController.DeleteUser)
}
