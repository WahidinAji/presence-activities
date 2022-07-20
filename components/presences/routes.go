package presences

import (
	"presence-activities/pkg/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func (d *PresenceDeps) PresenceRoutes(app *fiber.App) {
	api := app.Group("/api/presences", logger.New())
	api.Get("/", middleware.Protected(), d.GetAll)
	api.Post("/", middleware.Protected(), d.Presence)
}
