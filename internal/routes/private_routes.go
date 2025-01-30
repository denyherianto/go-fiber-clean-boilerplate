package routes

import (
	"github.com/denyherianto/go-fiber-clean-boilerplate/database"
	"github.com/denyherianto/go-fiber-clean-boilerplate/internal/middlewares"
	privateRoutes "github.com/denyherianto/go-fiber-clean-boilerplate/internal/routes/private_routes"
	"github.com/gofiber/fiber/v2"
)

// PrivateRoutes func for describe group of private routes.
func PrivateRoutes(a *fiber.App, db database.Database) {
	// API V1 Routes Group
	routeV1 := a.Group("/api/v1")
	routeV1.Use(middlewares.JWTProtected())
	routeV1.Use(middlewares.TokenValidation)

	privateRoutes.CompanyRoutes(routeV1, db)
}
