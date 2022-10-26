package routes

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"human-resources-backend/controllers"
	"human-resources-backend/middlewares"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RegisterCompanyRoutes(router fiber.Router, db *gorm.DB, awsSession *session.Session) {
	companyController := controllers.NewCompanyController(db, awsSession)

	company := router.Group("/companies")

	company.Get("/", middlewares.JWTProtected(), companyController.GetCompanyList)
	company.Get(":id", middlewares.JWTProtected(), companyController.GetCompanyById)
	company.Post("/", middlewares.JWTProtected(), companyController.CreateCompany)
	company.Put(":id", middlewares.JWTProtected(), companyController.UpdateCompany)
	company.Delete(":id", middlewares.JWTProtected(), companyController.DeleteCompany)
}
