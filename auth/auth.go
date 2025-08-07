package auth

import (
	"os"
	"time"

	"github.com/frank1nz/ManagementEmployee/database"
	"github.com/frank1nz/ManagementEmployee/models"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(bytes), err
}

func CheckPassword(hashed, plain string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(plain))
	return err == nil
}

func Signup(c *fiber.Ctx) error {
	var data models.Admin
	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	hashed, err := HashPassword(data.Password)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Failed to hash password"})
	}
	data.Password = hashed

	if err := database.DB.Create(&data).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "User already exists or invalid data"})
	}

	return c.JSON(fiber.Map{"message": "Signup successful"})

}

func Login(c *fiber.Ctx) error {
	var input models.Admin

	if err := c.BodyParser(&input); err != nil || input.Username == "" || input.Password == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid input",
		})
	}

	// ค้นหาผู้ใช้
	var admin models.Admin
	if err := database.DB.Where("username = ?", input.Username).First(&admin).Error; err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "User not found",
		})
	}

	// ตรวจ password
	if err := bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(input.Password)); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Invalid password",
		})
	}

	// สร้าง JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       admin.ID,
		"username": admin.Username,
		"role":     admin.Role,
		"exp":      time.Now().Add(time.Hour * 72).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Could not login",
		})
	}

	return c.JSON(fiber.Map{
		"token": tokenString,
	})

}

func Logout(c *fiber.Ctx) error {
	// ไม่มี session management เพราะ JWT เป็น stateless
	// แนะนำให้ client ฝั่ง frontend ลบ token ออกจาก local storage/cookie
	return c.JSON(fiber.Map{"message": "Logged out (client must delete token)"})
}
