package routes

import (
	"human-resources-backend/controllers"
	"human-resources-backend/middlewares"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RegisterEmployeeRoutes(router fiber.Router, db *gorm.DB) {
	employeeController := controllers.NewEmployeeController(db)

	employee := router.Group("/employees")

	employee.Get("/", middlewares.JWTProtected(), employeeController.GetEmployeeList)
	employee.Get(":id", middlewares.JWTProtected(), employeeController.GetEmployeeById)
	employee.Post("/", middlewares.JWTProtected(), employeeController.CreateEmployee)
	employee.Put(":id", middlewares.JWTProtected(), employeeController.UpdateEmployee)
	employee.Delete(":id", middlewares.JWTProtected(), employeeController.DeleteEmployee)
}
