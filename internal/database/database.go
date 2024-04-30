package database

import (
	"fmt"
	"go-bponline/m/internal/repositories"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func doMigration() {
	err := DB.AutoMigrate(&repositories.User{})
	if err != nil {
		log.Panic("Failed to migrate database", err.Error())
	}
}

func ConnectWithDatabase() *gorm.DB {
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable", dbHost, dbUser, dbPass, dbName, dbPort)

	var err error

	DB, err = gorm.Open(postgres.Open(dsn))

	if err != nil {
		log.Panic("failed to open connection to database. ", err.Error())
	}
	log.Print("Successfully connected to database")

	// Comment this for avoid the migration
	// doMigration()

	return DB
}
