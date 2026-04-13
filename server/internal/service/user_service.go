package service

import (
	"github.com/google/uuid"
	"github.com/xinbreak/movie-web-app/internal/models"
	"github.com/xinbreak/movie-web-app/internal/repository"
)

type UserService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) RegisterUser(user *models.User) error {
	return s.repo.Create(user)
}

func (s *UserService) GetUser(id uuid.UUID) (*models.User, error) {
	return s.repo.GetByID(id)
}

func (s *UserService) ListUsers() ([]models.User, error) {
	return s.repo.GetAll()
}
