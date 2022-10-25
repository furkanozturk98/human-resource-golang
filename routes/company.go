package routes

import (
	"human-resources-backend/controllers"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RegisterCompanyRoutes(router fiber.Router, db *gorm.DB) {
	companyController := controllers.NewCompanyController(db)

	Company := router.Group("/companies")

	Company.Get("/", companyController.GetCompanyList)
	Company.Get(":id", companyController.GetCompanyById)
	Company.Post("/", companyController.CreateCompany)
	Company.Put(":id", companyController.UpdateCompany)
	Company.Delete(":id", companyController.DeleteCompany)
}
