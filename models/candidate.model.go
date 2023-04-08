package models

import (
	"time"

	"github.com/google/uuid"
)

// CandidateUser struct represents the model for a Candidateuser
type CandidateUser struct {
	ID           uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	Name         string    `gorm:"type:varchar(255);not null"`
	Email        string    `gorm:"uniqueIndex;not null"`
	MatricNumber string    `gorm:"uniqueIndex;not null"`
	Password     string    `gorm:"not null"`
	Role         string    `gorm:"type:varchar(255);not null"`
	Verified     bool      `gorm:"not null"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

// CandidateSignUpInput represents info needed to sign up a Candidateuser
type CandidateSignUpInput struct {
	Name            string `json:"name" binding:"required"`
	Email           string `json:"email" binding:"required"`
	MatricNumber    string `json:"matric_number" binding:"required"`
	Password        string `json:"password" binding:"required,min=8"`
	PasswordConfirm string `json:"passwordConfirm" binding:"required"`
}

// CandidateSignInInput represents the struct for data that Candidateusers need to sign in
type CandidateSignInInput struct {
	MatricNumber string `json:"matric_number"  binding:"required"`
	Password     string `json:"password"  binding:"required"`
}

// CandidateUserResponse is the struct for the response returned to the Candidateuser
type CandidateUserResponse struct {
	ID           uuid.UUID `json:"id,omitempty"`
	Name         string    `json:"name,omitempty"`
	Email        string    `json:"email,omitempty"`
	MatricNumber string    `json:"matric_number,omitempty"`
	Role         string    `json:"role,omitempty"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
