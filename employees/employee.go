package employees

import (
	"github.com/frank1nz/ManagementEmployee/database"
	"github.com/frank1nz/ManagementEmployee/models"
	"github.com/gofiber/fiber/v2"
)

func GetAllEmployees(c *fiber.Ctx) error {
	var employees []models.Employee
	result := database.DB.Find(&employees)

	if result.Error != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retrieve employees",
		})
	}

	return c.JSON(employees)
}

func GetEmployeeByID(c *fiber.Ctx) error {
	id := c.Params("id")
	var employee models.Employee

	if err := database.DB.First(&employee, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Employee not found",
		})
	}

	return c.JSON(employee)
}

func CreateEmployee(c *fiber.Ctx) error {
	employee := new(models.Employee)
	if err := c.BodyParser(employee); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Failed to parse request body",
		})
	}
	database.DB.Create(employee)

	return c.JSON(employee)
}

func UpdateEmployee(c *fiber.Ctx) error {
	id := c.Params("id")
	var employee models.Employee

	if err := database.DB.First(&employee, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Employee not found",
		})
	}

	if err := c.BodyParser(&employee); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Failed to parse request body",
		})
	}

	database.DB.Save(&employee)
	return c.JSON(employee)
}

func DeleteEmployee(c *fiber.Ctx) error {
	id := c.Params("id")
	var employee models.Employee

	if err := database.DB.First(&employee, id).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Employee not found",
		})
	}

	database.DB.Delete(&employee)
	return c.SendStatus(fiber.StatusNoContent)
}
