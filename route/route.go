package route

import (
	"github.com/frank1nz/ManagementEmployee/auth"
	"github.com/frank1nz/ManagementEmployee/employees"
	"github.com/frank1nz/ManagementEmployee/middleware"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")

	})
	api := app.Group("/api")
	api.Post("/register", auth.Signup)
	api.Post("/login", auth.Login)
	api.Post("/logout", auth.Logout)

	api.Use(middleware.Protected())

	employeeGroup := api.Group("/employees")
	employeeGroup.Get("/", employees.GetAllEmployees)
	employeeGroup.Post("/", employees.CreateEmployee)
	employeeGroup.Get("/:id", employees.GetEmployeeByID)
	employeeGroup.Put("/:id", employees.UpdateEmployee)
	employeeGroup.Delete("/:id", employees.DeleteEmployee)

}
