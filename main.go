package main

import (
	"backend/cbt-backend/initializers"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

var (
	server *gin.Engine
)

// Here we load the envs with Viper and create a connection pool to Postgres DB
func init() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("? Could not load environment variables", err)
	}

	initializers.ConnectDB(&config)

	// Here we create Gin router and assign it to server variable
	server = gin.Default()
}

func main() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("? Could not load environment variables", err)
	}

	router := server.Group("/api")
	router.GET("/home", func(ctx *gin.Context) {
		message := "Welcome to my CBT app"
		ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": message})
	})

	log.Fatal(server.Run(":" + config.ServerPort))
}
