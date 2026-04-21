package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/xinbreak/movie-web-app/internal/models"
	dto "github.com/xinbreak/movie-web-app/internal/models/dtos"
	"github.com/xinbreak/movie-web-app/internal/services"
)

type UserController struct {
	svc *services.UserService
}

func NewUserController(svc *services.UserService) *UserController {
	return &UserController{svc: svc}
}

// CreateUser godoc
// @Summary Register a new user
// @Tags auth
// @Accept json
// @Produce json
// @Param input body dto.UserRegisterDTO true "Registration Data"
// @Success 201 {object} dto.UserResponseDTO
// @Failure 400 {object} map[string]string
// @Router /auth/register [post]
func (ctrl *UserController) CreateUser(c *gin.Context) {
	var input dto.UserRegisterDTO
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := models.User{
		Username: input.Username,
		Email:    input.Email,
		Password: input.Password,
	}

	if err := ctrl.svc.RegisterUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, dto.UserResponseDTO{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		AvatarURL: user.Avatar,
	})
}

// Login godoc
// @Summary User login
// @Tags auth
// @Accept json
// @Produce json
// @Param input body dto.UserLoginDTO true "Login Data"
// @Success 200 {object} dto.UserResponseDTO
// @Failure 401 {object} map[string]string
// @Router /auth/login [post]
func (ctrl *UserController) Login(c *gin.Context) {
	var input dto.UserLoginDTO
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := ctrl.svc.Login(input.Email, input.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid email or password"})
		return
	}

	c.JSON(http.StatusOK, dto.UserResponseDTO{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		AvatarURL: user.Avatar,
	})
}

// GetUser godoc
// @Summary Get user by ID
// @Tags users
// @Produce json
// @Param id path string true "User UUID"
// @Success 200 {object} dto.UserResponseDTO
// @Failure 404 {object} map[string]string
// @Router /users/{id} [get]
func (ctrl *UserController) GetUser(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid uuid"})
		return
	}

	user, err := ctrl.svc.GetUser(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	c.JSON(http.StatusOK, dto.UserResponseDTO{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		AvatarURL: user.Avatar,
	})
}

// GetUsers godoc
// @Summary List all users
// @Tags users
// @Produce json
// @Success 200 {array} dto.UserResponseDTO
// @Router /users [get]
func (ctrl *UserController) GetUsers(c *gin.Context) {
	users, err := ctrl.svc.ListUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := make([]dto.UserResponseDTO, len(users))
	for i, u := range users {
		response[i] = dto.UserResponseDTO{
			ID:        u.ID,
			Username:  u.Username,
			Email:     u.Email,
			AvatarURL: u.Avatar,
		}
	}

	c.JSON(http.StatusOK, response)
}

// UpdateUser godoc
// @Summary Update user profile
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "User UUID"
// @Param input body dto.UserUpdateDTO true "Update Data"
// @Success 200 {object} map[string]string
// @Router /users/{id} [put]
func (ctrl *UserController) UpdateUser(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid uuid"})
		return
	}

	var input dto.UserUpdateDTO
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := models.User{
		ID:       id,
		Username: input.Username,
		Avatar:   input.AvatarURL,
	}

	if err := ctrl.svc.UpdateUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "user updated"})
}

// DeleteUser godoc
// @Summary Delete user
// @Tags users
// @Param id path string true "User UUID"
// @Success 204 "No Content"
// @Router /users/{id} [delete]
func (ctrl *UserController) DeleteUser(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid uuid"})
		return
	}

	if err := ctrl.svc.DeleteUser(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}
