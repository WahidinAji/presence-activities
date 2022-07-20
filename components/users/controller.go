package users

import (
	"fmt"
	"presence-activities/pkg"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func (d *UserDeps) Login(c *fiber.Ctx) error {
	in := new(LoginIn)
	if err := c.BodyParser(in); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	resRepo, err := d.LoginRepo(c.Context(), *in)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Opps : " + err.Error()})
	}

	token, err := pkg.JWTToken(in.Email)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Generate token failed : " + err.Error()})
	}
	fmt.Println(resRepo.Email)

	res := new(JsonResponse)
	res.Email = resRepo.Email
	res.Token = token

	c.Cookie(&fiber.Cookie{
		Name:  "Authorization",
		
	})
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"data": res})
}

func (d *UserDeps) Register(c *fiber.Ctx) error {
	in := new(RegisIn)
	if err := c.BodyParser(in); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	bytes, err := bcrypt.GenerateFromPassword([]byte(in.Password), bcrypt.MinCost)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Hash password failed : " + err.Error()})
	}
	in.Password = string(bytes)

	resRepo, err := d.RegisterRepo(c.Context(), *in)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Opps : " + err.Error()})
	}

	token, err := pkg.JWTToken(in.Email)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Generate token failed : " + err.Error()})
	}

	// token := jwt.New(jwt.SigningMethodHS256)
	// claims := token.Claims.(jwt.MapClaims)
	// claims["email"] = resRepo.Email
	// claims["exp"] = time.Now().Add(time.Minute * 5).Unix()
	// jwt, err := token.SignedString([]byte(pkg.JWTSecret))
	// if err != nil {
	// 	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Generate token failed :" + err.Error()})
	// }

	res := new(JsonResponse)
	res.Email = resRepo.Email
	res.Token = token
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"data": res})
}
