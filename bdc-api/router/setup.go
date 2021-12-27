package router

import (
	"github.com/gofiber/fiber/v2"

	"go_server_test/controllers"
)

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	api.Get("/dashboard/:user_id", controllers.GetDashboardData)
	api.Post("/player", controllers.CreatePlayerAndJoinWorld)
}
