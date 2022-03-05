package main

import (
	"log"

	"go_server_test/controllers"
	"go_server_test/database"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func CreateServer() *fiber.App {
	app := fiber.New()
	return app
}

func main() {
	database.ConnectToDB()

	app := CreateServer()
	app.Use(cors.New(cors.Config{
		AllowHeaders:     "Access-Control-Allow-Headers, Origin, Content-Type, Accept, Authorization, Set-Cookie",
		AllowOrigins:     "*",
		AllowMethods:     "GET,HEAD,OPTIONS,POST,PUT,DELETE",
		AllowCredentials: true,
	}))
	controllers.SetupRoutes(app)

	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404)
	})

	log.Fatal(app.Listen(":3000"))
}
