package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Video struct {
	ID        uuid.UUID      `gorm:"type:uuid;primaryKey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	Title       string `gorm:"size:255;not null" json:"title"`
	Description string `gorm:"type:text" json:"description"`

	FilePath  string `gorm:"not null" json:"-"`
	Thumbnail string `json:"thumbnail_url"`

	Duration int    `json:"duration"`
	Views    uint32 `gorm:"default:0" json:"views_count"`

	UserID uuid.UUID `gorm:"type:uuid" json:"user_id"`

	Status string `gorm:"default:'uploaded'" json:"status"`
}

func (v *Video) BeforeCreate(tx *gorm.DB) (err error) {
	v.ID = uuid.New()
	return
}
