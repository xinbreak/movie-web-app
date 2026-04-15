package dto

import (
	"github.com/google/uuid"
)

type CreateCommentRequest struct {
	Content  string     `json:"content" binding:"required,min=1,max=5000"`
	PostID   uuid.UUID  `json:"post_id" binding:"required"`
	ParentID *uuid.UUID `json:"parent_id"`
}

type UpdateCommentRequest struct {
	Content string `json:"content" binding:"required,min=1,max=5000"`
}
