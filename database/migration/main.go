package main

import (
	"github.com/denyherianto/go-fiber-clean-boilerplate/database"
	"github.com/denyherianto/go-fiber-clean-boilerplate/internal/entities"
)

func main() {
	db := database.NewPostgresDatabase()

	db.GetDb().AutoMigrate(
		&entities.Company{},
	)
}
