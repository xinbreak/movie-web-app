package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID       uuid.UUID `gorm:"type:char(36);primaryKey" json:"id" swaggerignore:"true"`
	Username string    `gorm:"size:32;not null;uniqueIndex" json:"username"`
	Email    string    `gorm:"size:255;not null;uniqueIndex" json:"email"`
	Password string    `gorm:"not null" json:"password"`
	Avatar   string    `gorm:"size:512" json:"avatar_url"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.New()
	return
}
