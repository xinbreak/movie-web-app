package models

import (
	"github.com/google/uuid"
)

type User struct {
	ID       uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	Username string    `gorm:"size:32;not null;uniqueIndex" json:"username"`
	Email    string    `gorm:"size:255;not null;uniqueIndex" json:"email"`
	Password string    `gorm:"not null" json:"-"`
	Avatar   string    `gorm:"size:512" json:"avatar_url"`
}
