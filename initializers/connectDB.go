// This code connects the app to the Postgres server

package initializers

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DB exports the database connection from gorm
var DB *gorm.DB

// ConnectDB connects to the Postgres server
func ConnectDB(config *Config) {
	var err error

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Africa/Lagos", config.DBHost, config.DBUserName, config.DBUserPassword, config.DBName, config.DBPort)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the Database")
	}
	fmt.Println("🚀 Connected Successfully to the Database")
}
