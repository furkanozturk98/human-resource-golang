package controllers

import (
	"human-resources-backend/models"
	"human-resources-backend/validators"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type UserController struct {
	Database *gorm.DB
}

func NewUserController(database *gorm.DB) UserController {
	return UserController{
		Database: database,
	}
}

func (r *UserController) GetUserList(c *fiber.Ctx) error {
	userModel := models.NewUserModel(r.Database)

	users, err := userModel.GetUserList()

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
			"data": users,
		})
}

func (r *UserController) GetUserById(c *fiber.Ctx) error {
	userModel := models.NewUserModel(r.Database)

	id := c.Params("id")

	userId, err := strconv.Atoi(id)

	user, err := userModel.GetUserById(userId)

	if err != nil {
		return c.
			Status(fiber.StatusInternalServerError).
			JSON(fiber.Map{
				"error": err.Error(),
			})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"data": user})
}

func (r *UserController) CreateUser(c *fiber.Ctx) error {
	userModel := models.NewUserModel(r.Database)

	userValidator := new(validators.User)

	if err := c.BodyParser(userValidator); err != nil {
		return err
	}

	errors := validators.ValidateUser(*userValidator)

	if errors != nil {
		return c.
			Status(fiber.StatusBadRequest).
			JSON(errors)
	}

	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return err
	}

	err := userModel.CreateUser(&user)

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
			"data": fiber.Map{
				"FirstName": user.FirstName,
				"LastName":  user.LastName,
				"Email":     user.Email,
			},
		})

}

func (r *UserController) UpdateUser(c *fiber.Ctx) error {
	userModel := models.NewUserModel(r.Database)

	id := c.Params("id")

	userId, err := strconv.Atoi(id)

	user, err := userModel.GetUserById(userId)

	if err != nil {
		return c.
			Status(fiber.StatusInternalServerError).
			JSON(fiber.Map{
				"error": err.Error(),
			})
	}

	userValidator := new(validators.User)

	if err := c.BodyParser(userValidator); err != nil {
		return err
	}

	errors := validators.ValidateUser(*userValidator)

	if errors != nil {
		return c.
			Status(fiber.StatusBadRequest).
			JSON(errors)

	}

	var userBody models.User

	if err := c.BodyParser(&userBody); err != nil {
		return err
	}

	err = userModel.UpdateUser(&user, &userBody)

	if err != nil {
		return c.
			Status(fiber.StatusInternalServerError).
			JSON(fiber.Map{"error": err.Error()})
	}

	return c.
		Status(fiber.StatusOK).
		JSON(fiber.Map{
			"data": user,
		})

}

func (r *UserController) DeleteUser(c *fiber.Ctx) error {
	userModel := models.NewUserModel(r.Database)

	id := c.Params("id")

	userId, err := strconv.Atoi(id)

	_, err = userModel.GetUserById(userId)

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

	userModel.DeleteUser(userId)

	return c.Status(fiber.StatusOK).Send(nil)
}
