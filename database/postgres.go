package database

import (
	"log"
	"sync"

	"github.com/inditekno/office-backend/pkg"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type postgresDatabase struct {
	Db *gorm.DB
}

var (
	once       sync.Once
	dbInstance *postgresDatabase
)

func NewPostgresDatabase() Database {
	once.Do(func() {
		err := godotenv.Load()
		if err != nil {
			log.Println("Error loading .env file")
		}

		// Build PostgreSQL connection URL.
		dsn, err := pkg.ConnectionURLBuilder("postgres")
		if err != nil {
			log.Println(err)
		}

		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			panic("failed to connect database")
		}

		dbInstance = &postgresDatabase{Db: db}
	})

	return dbInstance
}

func (p *postgresDatabase) GetDb() *gorm.DB {
	return dbInstance.Db
}
