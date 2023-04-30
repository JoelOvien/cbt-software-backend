package models

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	ID         uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"_id"`
	Name       string    `gorm:"not null" gorm:"not null"json:"name" validate:"required,min=4,max=100"`
	Id_Number  string    `gorm:"not null" json:"id_number" validate:"required,min=4,max=100"`
	Password   string    `gorm:"not null" json:"password" validate:"required,min=8"`
	Email      string    `gorm:"not null" json:"email" validate:"email,required"`
	User_type  string    `gorm:"not null" json:"user_type" validate:"required,eq=ADMIN|eq=STUDENT|eq=EXAMINER"`
	Created_at time.Time `gorm:"not null" json:"created_at"`
	Updated_at time.Time `gorm:"not null" json:"updated_at"`
	User_id    string    `gorm:"not null" json:"user_id"`
}

type UserResponse struct {
	ID         uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"_id"`
	Name       string    `gorm:"not null" gorm:"not null"json:"name" validate:"required,min=4,max=100"`
	Id_Number  string    `gorm:"not null" json:"id_number" validate:"required,min=4,max=100"`
	Email      string    `gorm:"not null" json:"email" validate:"email,required"`
	User_type  string    `gorm:"not null" json:"user_type" validate:"required,eq=ADMIN|eq=STUDENT|eq=EXAMINER"`
	Created_at time.Time `gorm:"not null" json:"created_at"`
	Updated_at time.Time `gorm:"not null" json:"updated_at"`
	User_id    string    `gorm:"not null" json:"user_id"`
}
