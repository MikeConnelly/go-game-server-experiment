package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

var store = session.New()

func SetupRoutes(app *fiber.App) {
	api := app.Group("/api")

	api.Post("/login", Login)
	api.Get("/dashboard", GetDashboardData)
	api.Post("/player", CreatePlayerAndJoinWorld)
}
