package presences

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func (d *PresenceDeps) PresenceRoutes(app *fiber.App) {
	api := app.Group("/api/presences", logger.New())
	api.Get("/", d.GetAll)
	api.Post("/", d.Presence)
}
