package activities

import (
	"presence-activities/pkg/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func (d *ActivityDeps) ActivityRoutes(app *fiber.App) {
	api := app.Group("/api/activities", logger.New())
	api.Use(middleware.IsAuth)
	api.Get("/", d.GetAll)
	api.Post("/", d.Create)
	// api.Get("/", middleware.Protected(), d.GetAll)
	// api.Post("/", middleware.Protected(), d.Create)
}
