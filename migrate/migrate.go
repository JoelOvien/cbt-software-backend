package main

import (
	"backend/cbt-backend/models"
	"fmt"
	"log"

	"backend/cbt-backend/database"
)

// Here we load the environment variables and created the connection pool to Postgres db
func init() {
	config, err := database.LoadConfig(".")
	if err != nil {
		log.Fatal("? Could not load environment variables", err)
	}

	database.ConnectDB(&config)
}

func main() {
	// We evoke this function provided by GORM to create the db migration and push changes to the database
	database.DB.AutoMigrate(&models.User{})

	fmt.Println("🚀 Migration complete")
}
