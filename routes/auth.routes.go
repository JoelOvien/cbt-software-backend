package routes

import (
	"backend/cbt-backend/controllers"
	"backend/cbt-backend/middleware"

	"github.com/gin-gonic/gin"
)

// AuthRouteController defines our authController struct of type AdminAuthController
type AuthRouteController struct {
	authController controllers.AuthController
}

// NewAuthRouteController returns a new AuthRouteController
func NewAuthRouteController(authController controllers.AuthController) AuthRouteController {
	return AuthRouteController{authController}
}

// AuthRoute defines auth routes for admin login
func (rc *AuthRouteController) AuthRoute(rg *gin.RouterGroup) {
	router := rg.Group("/auth")

	router.POST("/register", rc.authController.SignUpUser)
	router.POST("/login", rc.authController.SignInUser)
	router.GET("/refresh", rc.authController.RefreshAccessToken)
	router.GET("/logout", middleware.DeserializeUser(), rc.authController.LogoutUser)
}
