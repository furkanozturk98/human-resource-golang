package routes

import (
	"human-resources-backend/controllers"
	"human-resources-backend/middlewares"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RegisterCompanyRoutes(router fiber.Router, db *gorm.DB) {
	companyController := controllers.NewCompanyController(db)

	Company := router.Group("/companies")

	Company.Get("/", middlewares.JWTProtected(), companyController.GetCompanyList)
	Company.Get(":id", middlewares.JWTProtected(), companyController.GetCompanyById)
	Company.Post("/", middlewares.JWTProtected(), companyController.CreateCompany)
	Company.Put(":id", middlewares.JWTProtected(), companyController.UpdateCompany)
	Company.Delete(":id", middlewares.JWTProtected(), companyController.DeleteCompany)
}
