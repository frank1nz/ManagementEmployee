package middleware

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func Protected() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// ดึง token จาก header
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Missing or malformed JWT"})
		}

		// แยก "Bearer tokenstring"
		tokenString := ""
		if len(authHeader) > 7 && authHeader[:7] == "Bearer " {
			tokenString = authHeader[7:]
		} else {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid token format"})
		}

		// ตรวจสอบ JWT token
		token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
			// ตรวจสอบว่า signing method ตรงกับที่คาดไว้
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fiber.ErrUnauthorized
			}
			// secret key ต้องเหมือนกับตอนสร้าง token
			return []byte(os.Getenv("JWT_SECRET")), nil

		})

		if err != nil || !token.Valid {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid or expired JWT"})
		}

		// ใส่ข้อมูล claims ลง context เพื่อใช้งานต่อ
		claims := token.Claims.(jwt.MapClaims)
		c.Locals("userID", claims["id"])
		c.Locals("username", claims["username"])
		c.Locals("role", claims["role"])

		return c.Next()
	}
}
