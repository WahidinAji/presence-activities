package middleware

import (
	// "errors"
	"presence-activities/pkg"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
	"github.com/golang-jwt/jwt/v4"
)

// Protected protect routes
func Protected() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey:   []byte(pkg.JWTSecret),
		ErrorHandler: jwtError,
	})
}

func jwtError(c *fiber.Ctx, err error) error {
	if err.Error() == "Missing or malformed JWT" {
		return c.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"status": "error", "message": "Missing or malformed JWT", "data": nil})
	}
	return c.Status(fiber.StatusUnauthorized).
		JSON(fiber.Map{"status": "error", "message": "Invalid or expired JWT", "data": nil})
}

func IsAuth(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")

	_, err := jwt.ParseWithClaims(cookie, &jwt.RegisteredClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(pkg.JWTSecret), nil
	})

	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  fiber.StatusUnauthorized,
			"message": "Unauthorized",
		})
	}

	return c.Next()
	// claims := token.Claims.(*jwt.RegisteredClaims)
	// fmt.Println(claims.VerifyExpiresAt())
	// token, err = jwt.Parse(cookie, func(t *jwt.Token) (interface{}, error) {
	// 	return []byte(pkg.JWTSecret), nil
	// })
	// if token.Valid {
	// 	return c.Next()
	// } else if errors.Is(err, jwt.ErrTokenMalformed) {
	// 	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
	// 		"status":  fiber.StatusUnauthorized,
	// 		"message": "Missing or malformed token",
	// 	})
	// } else if errors.Is(err, jwt.ErrTokenExpired) {
	// 	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
	// 		"status":  fiber.StatusUnauthorized,
	// 		"message": "Token expired",
	// 	})
	// } else {
	// 	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
	// 		"status":  fiber.StatusUnauthorized,
	// 		"message": "Invalid token",
	// 	})
	// }
}
