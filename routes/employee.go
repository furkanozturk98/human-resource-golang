package routes

import (
	"human-resources-backend/controllers"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RegisterEmployeeRoutes(router fiber.Router, db *gorm.DB) {
	employeeController := controllers.NewEmployeeController(db)

	employee := router.Group("/employees")

	employee.Get("/", employeeController.GetEmployeeList)
	employee.Get(":id", employeeController.GetEmployeeById)
	employee.Post("/", employeeController.CreateEmployee)
	employee.Put(":id", employeeController.UpdateEmployee)
	employee.Delete(":id", employeeController.DeleteEmployee)
}
