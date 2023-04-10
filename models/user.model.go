package models

import (
	"database/sql/driver"
	"fmt"
	"github.com/google/uuid"
	"time"
)

type userType string

const (
	Admin     userType = "ADMIN"
	Candidate userType = "CANDIDATE"
	Examiner  userType = "EXAMINER"
)

func (u userType) Value() (driver.Value, error) {
	return string(u), nil
}

func (u *userType) Scan(value interface{}) error {
	if value == nil {
		*u = ""
		return nil
	}
	switch s := value.(type) {
	case []byte:
		*u = userType(s)
		return nil
	case string:
		*u = userType(s)
		return nil
	default:
		return fmt.Errorf("unsupported Scan value type: %T", value)
	}
}

// User struct represents the model for a user
type User struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key" json:"id"`
	Name      string    `gorm:"type:varchar(255);not null" json:"name"`
	Email     string    `gorm:"not null" json:"email"`
	UserName  string    `gorm:"uniqueIndex; not null" json:"user_name"`
	Password  string    `gorm:"not null" json:"password"`
	UserType  userType  `gorm:"type:user_enum" json:"user_type"`
	CreatedAt time.Time `gorm:"not null" json:"created_at"`
	UpdatedAt time.Time `gorm:"not null" json:"updated_at"`
}
