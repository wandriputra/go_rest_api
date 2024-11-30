package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// User represents a user in the system with unique username and email.
// It includes the following fields:
// - ID: A unique identifier for the user.
// - Username: The username of the user, which must be unique and not null.
// - Email: The email address of the user, which must be unique and not null.
// - Password: The password of the user, which must not be null.
type User struct {
	gorm.Model

	ID       uuid.UUID `gorm:"type:uuid;"`
	Username string    `gorm:"unique;not null" json:"username"`
	Email    string    `gorm:"unique;not null" json:"email"`
	Password string    `gorm:"not null" json:"password"`
}

type Users struct {
	Users []User `json:"users"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New()
	return
}
