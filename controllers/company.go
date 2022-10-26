package controllers

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"human-resources-backend/models"
	"human-resources-backend/validators"
	"os"
	"strconv"
	"strings"
)

type CompanyController struct {
	Database   *gorm.DB
	AWSSession *session.Session
}

func NewCompanyController(database *gorm.DB, session *session.Session) CompanyController {
	return CompanyController{
		Database:   database,
		AWSSession: session,
	}
}

func (r *CompanyController) GetCompanyList(c *fiber.Ctx) error {

	companyModel := models.NewCompanyModel(r.Database)

	companies, err := companyModel.GetCompanyList()

	if err != nil {
		return c.
			Status(fiber.StatusInternalServerError).
			JSON(fiber.Map{
				"error": err.Error(),
			})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"data": companies})
}

func (r *CompanyController) GetCompanyById(c *fiber.Ctx) error {
	companyModel := models.NewCompanyModel(r.Database)

	id := c.Params("id")

	companyId, err := strconv.Atoi(id)

	company, err := companyModel.GetCompanyById(companyId)

	if err != nil {
		return c.
			Status(fiber.StatusInternalServerError).
			JSON(fiber.Map{
				"error": err.Error(),
			})
	}

	return c.
		Status(fiber.StatusOK).
		JSON(fiber.Map{
			"data": company,
		})
}

func (r *CompanyController) CreateCompany(c *fiber.Ctx) error {
	companyModel := models.NewCompanyModel(r.Database)

	companyValidator := new(validators.Company)
	if err := c.BodyParser(companyValidator); err != nil {
		return err
	}

	errors := validators.ValidateCompany(*companyValidator)
	if errors != nil {
		return c.
			Status(fiber.StatusBadRequest).
			JSON(errors)
	}

	fileHeader, err := c.FormFile("logo")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	headerType := fileHeader.Header["Content-Type"][0]
	if headerType != "" && !strings.HasPrefix(headerType, "image") {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Wrong file type."})
	}

	companyValidator.Logo = *fileHeader

	uploader := s3manager.NewUploader(r.AWSSession)
	bucket := os.Getenv("AWS_BUCKET_NAME")
	fmt.Print(bucket)
	file, err := fileHeader.Open()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	awsKey := uuid.NewString() + companyValidator.Logo.Filename

	upload, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucket),
		ACL:    aws.String("bucket-owner-full-control"),
		Key:    aws.String(awsKey),
		Body:   file,
	})
	if err != nil {
		return c.
			Status(fiber.StatusInternalServerError).
			JSON(fiber.Map{
				"error": "Failed to upload image. " + err.Error(),
			})
	}

	var company models.Company
	c.BodyParser(&company)

	/*err = companyModel.CreateCompany(&company)

	if err != nil {
		return c.
			Status(fiber.StatusInternalServerError).
			JSON(fiber.Map{
				"error": err.Error(),
			})
	}*/

	company.Logo = upload.Location

	if err := companyModel.CreateCompany(companyValidator, &company, awsKey, upload.Location); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.
		Status(fiber.StatusOK).
		JSON(fiber.Map{
			"data": company,
		})

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
