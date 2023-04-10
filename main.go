package main

import (
	"backend/cbt-backend/controllers"
	"backend/cbt-backend/initializers"
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
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("? Could not load environment variables", err)
	}

	initializers.ConnectDB(&config)

	AuthController = controllers.NewAuthController(initializers.DB)
	AuthRouteController = routes.NewAuthRouteController(AuthController)

	UserController = controllers.NewUserController(initializers.DB)
	UserRouteController = routes.NewRouteUserController(UserController)

	// Here we create Gin router and assign it to server variable
	server = gin.Default()
}

func main() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("? Could not load environment variables", err)
	}

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"http://localhost:8000", config.ClientOrigin}
	corsConfig.AllowCredentials = true

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
