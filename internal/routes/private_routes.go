package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/inditekno/office-backend/database"
	"github.com/inditekno/office-backend/internal/middlewares"
	privateRoutes "github.com/inditekno/office-backend/internal/routes/private_routes"
)

// PrivateRoutes func for describe group of private routes.
func PrivateRoutes(a *fiber.App, db database.Database) {
	// API V1 Routes Group
	routeV1 := a.Group("/api/v1")
	routeV1.Use(middlewares.JWTProtected())
	routeV1.Use(middlewares.TokenValidation)

	privateRoutes.CompanyRoutes(routeV1, db)
}
