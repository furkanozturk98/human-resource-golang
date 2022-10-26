package controllers

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"human-resources-backend/models"
	"human-resources-backend/services"
)

type AuthController struct {
	Database *gorm.DB
}

func NewAuthController(database *gorm.DB) AuthController {
	return AuthController{
		Database: database,
	}
}

type TokenRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (r *AuthController) Login(c *fiber.Ctx) error {
	userModel := models.NewUserModel(r.Database)

	var request TokenRequest
	var user models.User

	if err := c.BodyParser(&request); err != nil {
		return err
	}

	user, err := userModel.GetUserByEmail(request.Email)

	if err != nil {
		return c.
			Status(fiber.StatusInternalServerError).
			JSON(fiber.Map{
				"error": err.Error(),
			})
	}

	credentialError := userModel.CheckPassword(user.Password, request.Password)
	if credentialError != nil {
		return c.
			Status(fiber.StatusInternalServerError).
			JSON(fiber.Map{"error": err.Error()})

	}

	authService := services.NewAuthService(r.Database)

	tokenString, err := authService.GenerateJWT()

	if err != nil {
		return c.
			Status(fiber.StatusInternalServerError).
			JSON(fiber.Map{"error": err.Error()})
	}

	return c.
		Status(fiber.StatusOK).
		JSON(fiber.Map{"token": tokenString})
}

func (r *AuthController) GetNewAccessToken(c *fiber.Ctx) error {
	authService := services.NewAuthService(r.Database)

	token, err := authService.GenerateJWT()

	if err != nil {
		// Return status 500 and token generation error.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"error":        false,
		"msg":          nil,
		"access_token": token,
	})
}
