package users

import (
	"presence-activities/pkg/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func (d *UserDeps) UserRoutes(app *fiber.App) {
	// app.Use(csrf.New(csrf.Config{
	// 	KeyLookup:      "header:X-Csrf-Token",
	// 	CookieName:     "csrf_",
	// 	CookieSameSite: "Strict",
	// 	Expiration:     1 * time.Hour,
	// 	KeyGenerator:   utils.UUID,
	// }))
	api := app.Group("/api/users", logger.New())
	api.Post("/register", d.Register)
	api.Post("/login", d.Login)
	api.Post("/logout", middleware.IsAuth, d.SignOut)
}
