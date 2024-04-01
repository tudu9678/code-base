package model

import (
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// User ...
type User struct {
	ID          uuid.UUID  `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	FullName    string     `gorm:"not null"`
	PhoneNumber string     `gorm:"unique"`
	Email       string     `gorm:"unique"`
	UserName    string     `gorm:"unique"`
	Password    string     `gorm:"not null"`
	Dob         *time.Time `json:"birthday,omitempty"`
	LatestLogin *time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// TableName ...
func (User) TableName() string {
	return "users"
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	// You can add logic here to generate a secure hash for the password before saving
	// For example, using bcrypt:
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return
}
