package main

import (
	"backend/cbt-backend/controllers"
	"backend/cbt-backend/database"
	"backend/cbt-backend/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

var (
	server *gin.Engine
	// AuthController export
	AuthController controllers.AuthController
	// AuthRouteController export
	AuthRouteController routes.AuthRouteController
	// UserController export
	UserController controllers.UserController
	// UserRouteController export
	UserRouteController routes.UserRouteController
)

// Here we load the envs with Viper and create a connection pool to Postgres DB
func init() {
	config, err := database.LoadConfig(".")
	if err != nil {
		log.Fatal("? Could not load environment variables", err)
	}

	database.ConnectDB(&config)

	AuthController = controllers.NewAuthController(database.DB)
	AuthRouteController = routes.NewAuthRouteController(AuthController)

	UserController = controllers.NewUserController(database.DB)
	UserRouteController = routes.NewRouteUserController(UserController)

	// Here we create Gin router and assign it to server variable
	server = gin.Default()
}

func main() {
	config, err := database.LoadConfig(".")
	if err != nil {
		log.Fatal("? Could not load environment variables", err)
	}

	// Enable CORS
	server.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"http://localhost:8000", config.ClientOrigin}
	corsConfig.AllowCredentials = true
	corsConfig.AllowMethods = []string{"GET", "POST", "PUT", "DELETE"}            // Set the allowed HTTP methods
	corsConfig.AllowHeaders = []string{"Authorization", "Content-Type", "Accept"} // Set the allowed headers

	server.Use(cors.New(corsConfig))

	router := server.Group("/api")
	router.GET("/home", func(ctx *gin.Context) {
		message := "Welcome to my CBT app"
		ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": message})
	})

	AuthRouteController.AuthRoute(router)
	UserRouteController.UserRoute(router)
	log.Fatal(server.Run(":" + config.ServerPort))

	log.Fatal(server.Run(":" + config.ServerPort))
}
