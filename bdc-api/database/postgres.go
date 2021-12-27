package database

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	. "go_server_test/config"
)

var DB *gorm.DB

func ConnectToDB() {
	dsn := fmt.Sprintf("host=localhost user=%s password=%s dbname=%s port=%s sslmode=disable",
		Config("PSQL_USER"), Config("PSQL_PASS"), Config("PSQL_DBNAME"), Config("PSQL_PORT"))

	log.Print("Connecting to PostgreSQL DB... ")
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
		os.Exit(2)
	}

	log.Println("connected")
}
