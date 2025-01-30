package main

import (
	"github.com/inditekno/office-backend/cmd/server"
	"github.com/inditekno/office-backend/database"

	_ "github.com/inditekno/office-backend/docs" // load API Docs files (Swagger)

	_ "github.com/joho/godotenv/autoload" // load .env file automatically
)

// @title API
// @version 1.0
// @description This is an auto-generated API Docs.
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email your@mail.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @BasePath /api
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	db := database.NewPostgresDatabase()
	server.NewFiberServer(db).Start()
}
