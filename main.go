package main

import (
	"context"
	"log"
	"presence-activities/components/activities"
	"presence-activities/components/presences"
	"presence-activities/components/users"
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

	//users deps
	user := users.UserDeps{DB: conn}
	user.UserRoutes(app)

	//presences deps
	presence := presences.PresenceDeps{DB: conn}
	presence.PresenceRoutes(app)

	//activites deps
	activity := activities.ActivityDeps{DB: conn}
	activity.ActivityRoutes(app)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World! BY AJI")
	})

	app.Listen(":8080")
}
