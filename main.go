package main

import (
	"context"
	"fmt"
	"log"
	"presence-activities/components/activities"
	"presence-activities/components/presences"
	"presence-activities/components/users"
	"presence-activities/pkg"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
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

	//https
	
	//set cors
	app.Use(cors.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders:  "Origin, Content-Type, Accept,Bearer",
		AllowMethods:     "GET, POST, PATCH, PUT, DELETE",
	}))

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

	now := time.Now()
	fmt.Println(now.Format("2006-02-01"))
	app.Listen(":8080")
}
