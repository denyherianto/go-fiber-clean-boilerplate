package server

import (
	"os"

	"github.com/denyherianto/go-fiber-clean-boilerplate/configs"
	"github.com/denyherianto/go-fiber-clean-boilerplate/database"
	"github.com/denyherianto/go-fiber-clean-boilerplate/internal/middlewares"
	"github.com/denyherianto/go-fiber-clean-boilerplate/internal/routes"
	"github.com/denyherianto/go-fiber-clean-boilerplate/pkg"

	"github.com/gofiber/fiber/v2"

	_ "github.com/denyherianto/go-fiber-clean-boilerplate/docs" // load API Docs files (Swagger)

	_ "github.com/joho/godotenv/autoload" // load .env file automatically
)

type Server interface {
	Start()
}

type fiberServer struct {
	app *fiber.App
	db  database.Database
}

func NewFiberServer(db database.Database) Server {
	// Define Fiber config.
	config := configs.FiberConfig()
	// Define a new Fiber app with config.
	fiberApp := fiber.New(config)

	return &fiberServer{
		app: fiberApp,
		db:  db,
	}
}

func (s *fiberServer) Start() {
	// Middlewares.
	middlewares.FiberMiddleware(s.app) // Register Fiber's middleware for app.

	// Routes.
	routes.SwaggerRoute(s.app)        // Register a route for API Docs (Swagger).
	routes.PublicRoutes(s.app)        // Register a public routes for app.
	routes.PrivateRoutes(s.app, s.db) // Register a private routes for app.
	routes.NotFoundRoute(s.app)       // Register route for 404 Error.

	if os.Getenv("STAGE_STATUS") == "dev" {
		pkg.StartServer(s.app)
	} else {
		pkg.StartServerWithGracefulShutdown(s.app)
	}
}
