package dto

import "github.com/google/uuid"

type UserRegisterDTO struct {
	Username string `json:"username" binding:"required,min=3,max=32" example:"johndoe"`
	Email    string `json:"email" binding:"required,email" example:"john@example.com"`
	Password string `json:"password" binding:"required,min=8" example:"secret123"`
}

type UserLoginDTO struct {
	Email    string `json:"email" binding:"required,email" example:"john@example.com"`
	Password string `json:"password" binding:"required" example:"secret123"`
}

type UserResponseDTO struct {
	ID        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	AvatarURL string    `json:"avatar_url"`
}

type UserUpdateDTO struct {
	Username  string `json:"username" binding:"omitempty,min=3"`
	Password  string `json:"password" binding:"required" example:"secret123"`
	AvatarURL string `json:"avatar_url" binding:"omitempty,url"`
}
