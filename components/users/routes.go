package users

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func (d *UserDeps) UserRoutes(app *fiber.App) {
	api := app.Group("/api/users", logger.New())
	api.Post("/register", d.Register)
	api.Post("/login", d.Login)
}
