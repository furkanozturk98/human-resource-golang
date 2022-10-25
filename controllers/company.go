package controllers

import (
	"human-resources-backend/models"
	"human-resources-backend/validators"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type CompanyController struct {
	Database *gorm.DB
}

func NewCompanyController(database *gorm.DB) CompanyController {
	return CompanyController{
		Database: database,
	}
}

func (r *CompanyController) GetCompanyList(c *fiber.Ctx) error {

	CompanyModel := models.NewCompanyModel(r.Database)

	Companys, err := CompanyModel.GetCompanyList()

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"data": Companys})
}

func (r *CompanyController) GetCompanyById(c *fiber.Ctx) error {
	CompanyModel := models.NewCompanyModel(r.Database)

	id := c.Params("id")

	CompanyId, err := strconv.Atoi(id)

	Company, err := CompanyModel.GetCompanyById(CompanyId)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"data": Company})
}

func (r *CompanyController) CreateCompany(c *fiber.Ctx) error {
	CompanyModel := models.NewCompanyModel(r.Database)

	//var Company models.Company
	CompanyValidator := new(validators.Company)

	if err := c.BodyParser(CompanyValidator); err != nil {
		return err
	}

	errors := validators.ValidateCompany(*CompanyValidator)

	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)

	}

	var Company models.Company
	c.BodyParser(&Company)

	err := CompanyModel.CreateCompany(&Company)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.
		Status(fiber.StatusOK).
		JSON(fiber.Map{"data": Company})

}

func (r *CompanyController) UpdateCompany(c *fiber.Ctx) error {
	CompanyModel := models.NewCompanyModel(r.Database)

	id := c.Params("id")

	CompanyId, err := strconv.Atoi(id)

	Company, err := CompanyModel.GetCompanyById(CompanyId)

	if err != nil {
		return c.
			Status(fiber.StatusInternalServerError).
			JSON(fiber.Map{"error": err.Error()})
	}

	//var Company models.Company
	CompanyValidator := new(validators.Company)

	if err := c.BodyParser(CompanyValidator); err != nil {
		return err
	}

	errors := validators.ValidateCompany(*CompanyValidator)

	if errors != nil {
		return c.
			Status(fiber.StatusBadRequest).
			JSON(errors)

	}

	var CompanyBody models.Company
	c.BodyParser(&CompanyBody)

	err = CompanyModel.UpdateCompany(&Company, &CompanyBody)

	if err != nil {
		return c.
			Status(fiber.StatusInternalServerError).
			JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"data": Company})

}

func (r *CompanyController) DeleteCompany(c *fiber.Ctx) error {
	CompanyModel := models.NewCompanyModel(r.Database)

	id := c.Params("id")

	CompanyId, err := strconv.Atoi(id)

	_, err = CompanyModel.GetCompanyById(CompanyId)

	if err != nil {
		return c.
			Status(fiber.StatusInternalServerError).
			JSON(fiber.Map{"error": err.Error()})
	}

	if err != nil {
		return c.
			Status(fiber.StatusInternalServerError).
			JSON(fiber.Map{"error": err.Error()})
	}

	CompanyModel.DeleteCompany(CompanyId)

	return c.Status(fiber.StatusOK).Send(nil)

}
