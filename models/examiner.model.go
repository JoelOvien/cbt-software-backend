package models

import (
	"time"

	"github.com/google/uuid"
)

// ExaminerUser struct represents the model for a user
type ExaminerUser struct {
	ID          uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	Name        string    `gorm:"type:varchar(255);not null"`
	Email       string    `gorm:"uniqueIndex;not null"`
	StaffNumber string    `gorm:"uniqueIndex;not null"`
	Password    string    `gorm:"not null"`
	Role        string    `gorm:"type:varchar(255);not null"`
	Verified    bool      `gorm:"not null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// ExaminerSignUpInput represents info needed to sign up a user
type ExaminerSignUpInput struct {
	Name            string `json:"name" binding:"required"`
	Email           string `json:"email" binding:"required"`
	StaffNumber     string `json:"staff_number" binding:"required"`
	Password        string `json:"password" binding:"required,min=8"`
	PasswordConfirm string `json:"password_confirm" binding:"required"`
}

// ExaminerSignInInput represents the struct for data that users need to sign in
type ExaminerSignInInput struct {
	StaffNumber string `json:"staff_number"  binding:"required"`
	Password    string `json:"password"  binding:"required"`
}

// ExaminerUserResponse is the struct for the response returned to the user
type ExaminerUserResponse struct {
	ID          uuid.UUID `json:"id,omitempty"`
	Name        string    `json:"name,omitempty"`
	Email       string    `json:"email,omitempty"`
	StaffNumber string    `json:"staff_number,omitempty"`
	Role        string    `json:"role,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
