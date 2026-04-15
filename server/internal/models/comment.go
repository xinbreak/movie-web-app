package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Comment struct {
	ID        uuid.UUID  `gorm:"type:char(36);primaryKey" json:"id" swaggerignore:"true"`
	Content   string     `gorm:"type:text;not null" json:"content"`
	CreatedAt time.Time  `gorm:"autoCreateTime" json:"created_at"`
	UserID    uuid.UUID  `gorm:"type:char(36);not null;index" json:"user_id"`
	ParentID  *uuid.UUID `gorm:"type:char(36);index" json:"parent_id,omitempty"`
	// VideoID   uuid.UUID `gorm:"type:char(36);not null;index" json:"post_id"`

	User    User      `gorm:"foreignKey:UserID" json:"user,omitempty"`
	Parent  *Comment  `gorm:"foreignKey:ParentID" json:"parent,omitempty"`
	Replies []Comment `gorm:"foreignKey:ParentID" json:"replies,omitempty"`
	// Video Video `gorm:"foreignKey:PostID" json:"post,omitempty"`
}

func (c *Comment) BeforeCreate(tx *gorm.DB) (err error) {
	c.ID = uuid.New()
	return
}

func (c *Comment) IsRootComment() bool {
	return c.ParentID == nil
}

func (c *Comment) GetRepliesCount(db *gorm.DB) (int64, error) {
	var count int64
	err := db.Model(&Comment{}).Where("parent_id = ?", c.ID).Count(&count).Error
	return count, err
}

func (c *Comment) GetCommentWithUser(db *gorm.DB) error {
	return db.Preload("User").First(c, c.ID).Error
}
