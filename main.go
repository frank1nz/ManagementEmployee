package main

import (
	"github.com/frank1nz/ManagementEmployee/database"
	"github.com/frank1nz/ManagementEmployee/models"
	"github.com/frank1nz/ManagementEmployee/route"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	app := fiber.New()
	database.Connect()

	database.DB.AutoMigrate(&models.Employee{})
	database.DB.AutoMigrate(&models.Admin{})
	route.SetupRoutes(app)
	app.Listen(":8080")

}
