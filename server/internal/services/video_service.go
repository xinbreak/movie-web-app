package services

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/xinbreak/movie-web-app/internal/models"
	"github.com/xinbreak/movie-web-app/internal/repositories"
)

var (
	ErrVideoNotFound     = errors.New("video not found")
	ErrVideoEmptyTitle   = errors.New("video title cannot be empty")
	ErrUnauthorizedVideo = errors.New("you are not authorized to perform this action")
	ErrInvalidVideoID    = errors.New("invalid video id")
)

type VideoService interface {
	CreateVideo(video *models.Video) (*models.Video, error)
	GetVideoByID(id uuid.UUID) (*models.Video, error)
	GetAllVideos(page, pageSize int) ([]models.Video, int64, error)
	GetVideosByUser(userID uuid.UUID, page, pageSize int) ([]models.Video, int64, error)
	UpdateVideo(id, userID uuid.UUID, title, description string) (*models.Video, error)
	DeleteVideo(id, userID uuid.UUID) error
	IncrementViews(id uuid.UUID) error
}

type videoService struct {
	videoRepo repositories.VideoRepository
}

func NewVideoService(videoRepo repositories.VideoRepository) VideoService {
	return &videoService{
		videoRepo: videoRepo,
	}
}

func (s *videoService) CreateVideo(video *models.Video) (*models.Video, error) {
	if video.Title == "" {
		return nil, ErrVideoEmptyTitle
	}

	if err := s.videoRepo.Create(video); err != nil {
		return nil, fmt.Errorf("failed to create video: %w", err)
	}

	created, err := s.videoRepo.GetByID(video.ID)
	if err != nil {
		return nil, err
	}

	return created, nil
}

func (s *videoService) GetVideoByID(id uuid.UUID) (*models.Video, error) {
	video, err := s.videoRepo.GetByID(id)
	if err != nil {
		return nil, err
	}
	if video == nil {
		return nil, ErrVideoNotFound
	}
	return video, nil
}

func (s *videoService) GetAllVideos(page, pageSize int) ([]models.Video, int64, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 20
	}

	videos, total, err := s.videoRepo.GetAll(page, pageSize)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to get videos: %w", err)
	}

	return videos, total, nil
}

func (s *videoService) GetVideosByUser(userID uuid.UUID, page, pageSize int) ([]models.Video, int64, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 20
	}

	videos, total, err := s.videoRepo.GetByUserID(userID, page, pageSize)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to get user videos: %w", err)
	}

	return videos, total, nil
}

func (s *videoService) UpdateVideo(id, userID uuid.UUID, title, description string) (*models.Video, error) {
	video, err := s.videoRepo.GetByID(id)
	if err != nil {
		return nil, err
	}
	if video == nil {
		return nil, ErrVideoNotFound
	}

	isOwner, err := s.videoRepo.IsOwner(id, userID)
	if err != nil {
		return nil, err
	}
	if !isOwner {
		return nil, ErrUnauthorizedVideo
	}

	if title == "" {
		return nil, ErrVideoEmptyTitle
	}

	video.Title = title
	video.Description = description

	if err := s.videoRepo.Update(video); err != nil {
		return nil, fmt.Errorf("failed to update video: %w", err)
	}

	return video, nil
}

func (s *videoService) DeleteVideo(id, userID uuid.UUID) error {
	exists, err := s.videoRepo.ExistsByID(id)
	if err != nil {
		return err
	}
	if !exists {
		return ErrVideoNotFound
	}

	isOwner, err := s.videoRepo.IsOwner(id, userID)
	if err != nil {
		return err
	}
	if !isOwner {
		return ErrUnauthorizedVideo
	}

	if err := s.videoRepo.Delete(id); err != nil {
		return fmt.Errorf("failed to delete video: %w", err)
	}

	return nil
}

func (s *videoService) IncrementViews(id uuid.UUID) error {
	exists, err := s.videoRepo.ExistsByID(id)
	if err != nil {
		return err
	}
	if !exists {
		return ErrVideoNotFound
	}

	return s.videoRepo.IncrementViews(id)
}
