package main

import (
	"backend/cbt-backend/models"
	"fmt"
	"log"

	"backend/cbt-backend/initializers"
)

// Here we load the environment variables and created the connection pool to Postgres db
func init() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("? Could not load environment variables", err)
	}

	initializers.ConnectDB(&config)
}

func main() {
	// We evoke this function provided by GORM to create the db migration and push changes to the database
	initializers.DB.AutoMigrate(&models.AdminUser{})
	initializers.DB.AutoMigrate(&models.CandidateUser{})
	initializers.DB.AutoMigrate(&models.ExaminerUser{})

	fmt.Println("🚀 Migration complete")
}
