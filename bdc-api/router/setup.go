package router

import (
	"github.com/gofiber/fiber/v2"

	"go_server_test/handlers"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	api.Get("/dashboard/:user_id", handlers.GetDashboardData)
	api.Post("/player", handlers.CreatePlayerAndJoinWorld)
}
