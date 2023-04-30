package controllers

import (
	"backend/cbt-backend/database"
	"backend/cbt-backend/models"
	"backend/cbt-backend/utils"
	"github.com/google/uuid"

	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"

	"gorm.io/gorm"
)

// AuthController for this
type AuthController struct {
	DB *gorm.DB
}

// NewAuthController for this
func NewAuthController(DB *gorm.DB) AuthController {
	return AuthController{DB}
}

// SignUpUser SignUp User
func (aac *AuthController) SignUpUser(ctx *gin.Context) {
	var payload *models.User

	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": err.Error()})
		return
	}

	if payload.Name == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "name field is required"})
		return
	}
	if payload.Email == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "email field is required"})
		return
	}
	if payload.Id_Number == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "id_number field is required"})
		return
	}

	if payload.User_type == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "user_type field is required"})
		return
	}

	if payload.Password == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "password field is required"})
		return
	}

	hashedPassword, err := utils.HashPassword(payload.Password)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": err.Error()})
		return
	}
	// First, check if a user with the same Id_Number already exists
	existingUser := &models.User{}
	existingIDResult := aac.DB.Where("id_number = ?", payload.Id_Number).First(existingUser)
	if existingIDResult.Error == nil {
		ctx.JSON(http.StatusConflict, gin.H{"status": http.StatusBadRequest, "message": "User with that Id already exists"})
		return
	}

	existingEmailResult := aac.DB.Where("email = ?", payload.Email).First(existingUser)
	if existingEmailResult.Error == nil {
		ctx.JSON(http.StatusConflict, gin.H{"status": http.StatusBadRequest, "message": "User with that email already exists"})
		return
	}

	now := time.Now()
	newUser := models.User{
		Name:       payload.Name,
		Id_Number:  payload.Id_Number,
		Password:   hashedPassword,
		Email:      strings.ToLower(payload.Email),
		User_type:  payload.User_type,
		Created_at: now,
		Updated_at: now,
		User_id:    payload.Id_Number + "_" + payload.User_type,
	}

	result := aac.DB.Create(&newUser)

	if result.Error != nil && strings.Contains(result.Error.Error(), "duplicate key value violates unique") {
		ctx.JSON(http.StatusConflict, gin.H{"status": http.StatusBadRequest, "message": "User with that Id already exists"})
		return
	} else if result.Error != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": "Something bad happened"})
		return
	}

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Status":  http.StatusBadRequest,
			"Message": "error",
			"Data":    map[string]interface{}{"data": err.Error()}})
		return
	}

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Status":  http.StatusInternalServerError,
			"Message": "error",
			"Data":    map[string]interface{}{"data": err.Error()}})
		return
	}

	userResponse := &models.UserResponse{
		ID:         newUser.ID,
		Name:       newUser.Name,
		Email:      newUser.Email,
		Id_Number:  newUser.Id_Number,
		User_type:  newUser.User_type,
		Created_at: newUser.Created_at,
		Updated_at: newUser.Updated_at,
		User_id:    newUser.User_id,
	}

	ctx.JSON(http.StatusCreated, gin.H{"status": "success", "data": gin.H{"user": userResponse}})
}

// SignInUser for user sign in
func (aac *AuthController) SignInUser(ctx *gin.Context) {
	var payload *models.User
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": err.Error()})
		return
	}

	var user models.User
	result := aac.DB.First(&user, "Id_Number = ?", payload.Id_Number)

	if result.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Invalid id or Password"})
		return
	}

	if user.ID == uuid.Nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Invalid id or Password"})
		return
	}

	if err := utils.VerifyPassword(user.Password, payload.Password); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": "Invalid Password"})
		return
	}

	config, _ := database.LoadConfig(".")

	// Generate Tokens
	accessToken, err := utils.CreateToken(config.AccessTokenExpiresIn, user.ID, config.AccessTokenPrivateKey)
	if err != nil {
		fmt.Printf("failed to create token")
		ctx.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": err.Error()})
		return
	}

	refreshToken, err := utils.CreateToken(config.RefreshTokenExpiresIn, user.ID, config.RefreshTokenPrivateKey)
	if err != nil {
		fmt.Printf("failed to create refresh token")

		ctx.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": err.Error()})
		return
	}

	ctx.SetCookie("access_token", accessToken, config.AccessTokenMaxAge*60, "/", "localhost", false, true)
	ctx.SetCookie("refresh_token", refreshToken, config.RefreshTokenMaxAge*60, "/", "localhost", false, true)
	ctx.SetCookie("logged_in", "true", config.AccessTokenMaxAge*60, "/", "localhost", false, false)

	userResponse := &models.UserResponse{
		ID:         user.ID,
		Name:       user.Name,
		Email:      user.Email,
		Id_Number:  user.Id_Number,
		User_type:  user.User_type,
		Created_at: user.Created_at,
		Updated_at: user.Updated_at,
		User_id:    user.User_id,
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "access_token": accessToken, "refresh_token": refreshToken, "user": userResponse})
}

// RefreshAccessToken to refresh Admin access toke
func (aac *AuthController) RefreshAccessToken(ctx *gin.Context) {
	message := "could not refresh access token"

	cookie, err := ctx.Cookie("refresh_token")

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"status": "fail", "message": message})
		return
	}

	config, _ := database.LoadConfig(".")

	sub, err := utils.ValidateToken(cookie, config.RefreshTokenPublicKey)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	var user models.User
	result := aac.DB.First(&user, "id = ?", fmt.Sprint(sub))
	if result.Error != nil {
		ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"status": "fail", "message": "the user belonging to this token no logger exists"})
		return
	}

	accessToken, err := utils.CreateToken(config.AccessTokenExpiresIn, user.ID, config.AccessTokenPrivateKey)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"status": "fail", "message": err.Error()})
		return
	}

	ctx.SetCookie("access_token", accessToken, config.AccessTokenMaxAge*60, "/", "localhost", false, true)
	ctx.SetCookie("logged_in", "true", config.AccessTokenMaxAge*60, "/", "localhost", false, false)

	ctx.JSON(http.StatusOK, gin.H{"status": "success", "access_token": accessToken})
}

// LogoutUser to log out
func (aac *AuthController) LogoutUser(ctx *gin.Context) {
	ctx.SetCookie("access_token", "", -1, "/", "localhost", false, true)
	ctx.SetCookie("refresh_token", "", -1, "/", "localhost", false, true)
	ctx.SetCookie("logged_in", "", -1, "/", "localhost", false, false)

	ctx.JSON(http.StatusOK, gin.H{"status": "success"})
}
