package routes

import (
	"github.com/gofiber/fiber/v2"
)

// PublicRoutes func for describe group of public routes.
func PublicRoutes(a *fiber.App) {
	// Health Check
	a.Get("/healthz", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})
	a.Get("/test-office", func(c *fiber.Ctx) error {
		return c.SendString("office OK")
	})
}
