package middleware

import (
	"github.com/golang-jwt/jwt/v4/request"

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

	token, err := jwt.ParseWithClaims(cookie, &jwt.RegisteredClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(pkg.JWTSecret), nil
	})
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  fiber.StatusUnauthorized,
			"message": "Unauthenticated",
			//"message": "Unauthorized",
		})
	}

	//check if the token is nil
	if token == nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  fiber.StatusUnauthorized,
			"message": "Unauthorized token nil",
		})
	}
	//if claims, ok := token.Claims.(*jwt.RegisteredClaims); ok && token.Valid {
	//	fmt.Printf("%v %v", claims, claims.Issuer)
	//} else {
	//	fmt.Println(err)
	//}

	//check Authorization bearer token
	bearer := c.Get("Authorization")
	//check if bearer token is empty
	if bearer == "" {
		return c.Status(fiber.StatusNonAuthoritativeInformation).JSON(fiber.Map{
			"status":  fiber.StatusNonAuthoritativeInformation,
			"message": request.ErrNoTokenInRequest,
		})
	}
	//check if bearer not equal with token.Raw that's already on the cookies server
	if bearer != "Bearer "+token.Raw {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"status":  fiber.StatusUnauthorized,
			"message": "Unauthorized bearer",
		})
	}

	//jwtToken := c.Get("token")
	//fmt.Println(token.Claims.(*jwt.RegisteredClaims), token.Valid, token.Method, jwtToken, c.Get("Bearer"))
	//fmt.Println(token.Raw)
	//fmt.Println(cookie)
	//fmt.Println(bearer)

	//err = checkBearerToken(c, token.Raw)
	//if err != nil {
	//	return err
	//}
	//c.Get("Authorization","Bearer")

	//#region
	//token, err = jwt.Parse(cookie, func(t *jwt.Token) (interface{}, error) {
	//	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
	//		return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
	//	}
	//	return []byte(pkg.JWTSecret), nil
	//})
	//if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
	//	fmt.Println(claims["email"], claims["exp"], token.Header["token"], token.Method)
	//} else {
	//	fmt.Println(err)
	//}
	//#endregion

	//if claims, ok := token.Claims.(*MyCustomClaims); ok && token.Valid {
	//	fmt.Printf("%v %v", claims.Foo, claims.RegisteredClaims.Issuer)
	//} else {
	//	fmt.Println(err)
	//}

	//if all checked is accepted, user can access to their request
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
