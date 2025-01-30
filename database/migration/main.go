package main

import (
	"github.com/inditekno/office-backend/database"
	"github.com/inditekno/office-backend/internal/entities"
)

func main() {
	db := database.NewPostgresDatabase()

	db.GetDb().AutoMigrate(
		&entities.Company{},
	)
}
