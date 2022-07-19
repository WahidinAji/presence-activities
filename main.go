package main

import (
	"context"
	"log"
	"presence-activities/components/presences"
	"presence-activities/pkg"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v4"
)

func main() {
	app := fiber.New()

	ctx := context.Background()

	conn, err := pgx.Connect(ctx, pkg.PG_URL)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close(ctx)

	err = pkg.Migrate(ctx, conn)
	if err != nil {
		log.Fatal(err)
	}

	presence := presences.PresenceDeps{DB: conn}
	presence.PresenceRoutes(app)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	log.Fatal(app.Listen(":3000"))
}
