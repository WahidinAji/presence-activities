package presences

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func (d *PresenceDeps) GetAll(c *fiber.Ctx) error {
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

func (d *PresenceDeps) Presence(c *fiber.Ctx) error {
	in := new(PresenceIn)
	if err := c.BodyParser(in); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
			"error": err.Error(),
		})
	}

	res, err := d.PresenceRepo(c.Context(), *in)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(&fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"data": res})
}
