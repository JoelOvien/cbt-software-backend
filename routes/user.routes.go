package routes

import (
	"backend/cbt-backend/controllers"
	"backend/cbt-backend/middleware"
	"github.com/gin-gonic/gin"
)

// UserRouteController for
type UserRouteController struct {
	UserController controllers.UserController
}

// NewRouteUserController for
func NewRouteUserController(userController controllers.UserController) UserRouteController {
	return UserRouteController{userController}
}

// UserRoute for s
func (uc *UserRouteController) UserRoute(rg *gin.RouterGroup) {

	router := rg.Group("users")
	router.GET("/me", middleware.DeserializeUser(), uc.UserController.GetMe)
}
