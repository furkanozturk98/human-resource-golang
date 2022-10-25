package controllers

import (
	"human-resources-backend/models"
	"human-resources-backend/validators"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type EmployeeController struct {
	Database *gorm.DB
}

func NewEmployeeController(database *gorm.DB) EmployeeController {
	return EmployeeController{
		Database: database,
	}
}

func (r *EmployeeController) GetEmployeeList(c *fiber.Ctx) error {

	employeeModel := models.NewEmployeeModel(r.Database)

	employees, err := employeeModel.GetEmployeeList()

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"data": employees})
}

func (r *EmployeeController) GetEmployeeById(c *fiber.Ctx) error {
	employeeModel := models.NewEmployeeModel(r.Database)

	id := c.Params("id")

	EmployeeId, err := strconv.Atoi(id)

	employee, err := employeeModel.GetEmployeeById(EmployeeId)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"data": employee})
}

func (r *EmployeeController) CreateEmployee(c *fiber.Ctx) error {
	employeeModel := models.NewEmployeeModel(r.Database)

	//var Employee models.Employee
	employeeValidator := new(validators.Employee)

	if err := c.BodyParser(employeeValidator); err != nil {
		return err
	}

	errors := validators.ValidateEmployee(*employeeValidator)

	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(errors)

	}

	var Employee models.Employee
	c.BodyParser(&Employee)

	err := employeeModel.CreateEmployee(&Employee)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.
		Status(fiber.StatusOK).
		JSON(fiber.Map{"data": Employee})

}

func (r *EmployeeController) UpdateEmployee(c *fiber.Ctx) error {
	employeeModel := models.NewEmployeeModel(r.Database)

	id := c.Params("id")

	EmployeeId, err := strconv.Atoi(id)

	Employee, err := employeeModel.GetEmployeeById(EmployeeId)

	if err != nil {
		return c.
			Status(fiber.StatusInternalServerError).
			JSON(fiber.Map{"error": err.Error()})
	}

	//var Employee models.Employee
	employeeValidator := new(validators.Employee)

	if err := c.BodyParser(employeeValidator); err != nil {
		return err
	}

	errors := validators.ValidateEmployee(*employeeValidator)

	if errors != nil {
		return c.
			Status(fiber.StatusBadRequest).
			JSON(errors)

	}

	var EmployeeBody models.Employee
	c.BodyParser(&EmployeeBody)

	err = employeeModel.UpdateEmployee(&Employee, &EmployeeBody)

	if err != nil {
		return c.
			Status(fiber.StatusInternalServerError).
			JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"data": Employee})

}

func (r *EmployeeController) DeleteEmployee(c *fiber.Ctx) error {
	employeeModel := models.NewEmployeeModel(r.Database)

	id := c.Params("id")

	EmployeeId, err := strconv.Atoi(id)

	_, err = employeeModel.GetEmployeeById(EmployeeId)

	if err != nil {
		return c.
			Status(fiber.StatusInternalServerError).
			JSON(fiber.Map{"error": err.Error()})
	}

	err = employeeModel.DeleteEmployee(EmployeeId)

	if err != nil {
		return c.
			Status(fiber.StatusInternalServerError).
			JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).Send(nil)

}
