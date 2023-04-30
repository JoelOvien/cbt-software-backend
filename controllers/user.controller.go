package controllers

import (
	"backend/cbt-backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// UserController struct with type gormDB
type UserController struct {
	DB *gorm.DB
}

// NewUserController returns UserController with DB passed as it's value
func NewUserController(DB *gorm.DB) UserController {
	return UserController{DB}
}

// GetMe function for getting current admin user
func (uc *UserController) GetMe(ctx *gin.Context) {
	currentUser := ctx.MustGet("currentUser").(models.User)

	userResponse := &models.User{
		ID:         currentUser.ID,
		Name:       currentUser.Name,
		Id_Number:  currentUser.Id_Number,
		Email:      currentUser.Email,
		User_type:  currentUser.User_type,
		Created_at: currentUser.Created_at,
		Updated_at: currentUser.Updated_at,
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "data": gin.H{"user": userResponse}})
}
