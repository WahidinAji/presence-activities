package users

import (
	"fmt"
	"presence-activities/pkg"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"golang.org/x/crypto/bcrypt"
)

var store = session.New()

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

	// sess, err := store.Get(c)
	// if err != nil {
	// 	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	// }
	// sess.Set("authenticate", token)
	// // Save session
	// // if err := sess.Save(); err != nil {
	// // 	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	// // }
	// fmt.Println(sess.Keys(), sess.Get("authenticate"))

	cookie := fiber.Cookie{
		Name:  "jwt",
		Value: token,
		// Expires:  time.Now().Add(time.Hour * 24),
		Expires:  time.Now().Add(time.Minute * 2),
		HTTPOnly: true,
	}
	c.Cookie(&cookie)

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

	res := new(JsonResponse)
	res.Email = resRepo.Email
	res.Token = token

	// cookie := fiber.Cookie{
	// 	Name:    "jwt",
	// 	Value:   token,
	// 	Expires: time.Now().Add(time.Hour * 24),
	// 	HTTPOnly: true,
	// }
	// c.Cookie(&cookie)

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"data": res})
}

func (d *UserDeps) SignOut(c *fiber.Ctx) error {
	in := new(LoginOut)
	if err := c.BodyParser(in); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	res, err := d.SignOutRepo(c.Context(), *in)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Opps : " + err.Error()})
	}

	// token, err := pkg.JWTToken(in.Email)
	// if err != nil {
	// 	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Generate token failed : " + err.Error()})
	// }
	// c.ClearCookie("jwt", "session")
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}
	c.Cookie(&cookie)
	// sess, err := store.Get(c)
	// if err != nil {
	// 	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Opps : " + err.Error()})
	// }
	// // Delete key
	// sess.Delete("authenticate")

	// // Destroy session
	// if err := sess.Destroy(); err != nil {
	// 	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Opps : " + err.Error()})
	// }
	// fmt.Println(sess.Keys(), sess.Get("authenticate"))
	// c.ClearCookie()
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"data": nil, "message": res})
}
