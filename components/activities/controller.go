package activities

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func (d *ActivityDeps) Create(c *fiber.Ctx) error {
	// userId := c.FormValue("user_id")
	// if userId == "" {
	// 	return c.Status(400).SendString("Missing user_id")
	// }


	in := new(ActivityIn)
	if err := c.BodyParser(in); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	res, err := d.CreateRepo(c.Context(), *in)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"error": err.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"data": res})
}

func (d *ActivityDeps) GetAll(c *fiber.Ctx) error {
	listIn := new(ListIn)
	if err := c.BodyParser(listIn); err != nil {
		return c.Status(500).JSON(fiber.Map{"error": fmt.Sprint("lintIn", err.Error())})
	}

	res, err := d.FindAll(c.Context(), *listIn)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(200).JSON(fiber.Map{"data": res})
}

// func (d *ActivityDeps) GetByDate(ctx *fiber.Ctx) error {
	
// }