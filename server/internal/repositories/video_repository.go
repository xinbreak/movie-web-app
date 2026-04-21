package repositories

import (
	"errors"

	"github.com/google/uuid"
	"github.com/xinbreak/movie-web-app/internal/models"
	"gorm.io/gorm"
)

type VideoRepository interface {
	Create(video *models.Video) error
	GetByID(id uuid.UUID) (*models.Video, error)
	GetAll(page, pageSize int) ([]models.Video, int64, error)
	GetByUserID(userID uuid.UUID, page, pageSize int) ([]models.Video, int64, error)
	Update(video *models.Video) error
	Delete(id uuid.UUID) error
	IncrementViews(id uuid.UUID) error
	ExistsByID(id uuid.UUID) (bool, error)
	IsOwner(videoID, userID uuid.UUID) (bool, error)
}

type videoRepository struct {
	db *gorm.DB
}

func NewVideoRepository(db *gorm.DB) VideoRepository {
	return &videoRepository{db: db}
}

func (r *videoRepository) Create(video *models.Video) error {
	if video.ID == uuid.Nil {
		video.ID = uuid.New()
	}
	return r.db.Create(video).Error
}

func (r *videoRepository) GetByID(id uuid.UUID) (*models.Video, error) {
	var video models.Video
	err := r.db.Preload("User").First(&video, "id = ?", id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &video, nil
}

func (r *videoRepository) GetAll(page, pageSize int) ([]models.Video, int64, error) {
	var videos []models.Video
	var total int64

	if err := r.db.Model(&models.Video{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * pageSize
	err := r.db.Preload("User").
		Order("created_at DESC").
		Offset(offset).
		Limit(pageSize).
		Find(&videos).Error

	if err != nil {
		return nil, 0, err
	}

	return videos, total, nil
}

func (r *videoRepository) GetByUserID(userID uuid.UUID, page, pageSize int) ([]models.Video, int64, error) {
	var videos []models.Video
	var total int64

	query := r.db.Model(&models.Video{}).Where("user_id = ?", userID)

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * pageSize
	err := r.db.Preload("User").
		Where("user_id = ?", userID).
		Order("created_at DESC").
		Offset(offset).
		Limit(pageSize).
		Find(&videos).Error

	if err != nil {
		return nil, 0, err
	}

	return videos, total, nil
}

func (r *videoRepository) Update(video *models.Video) error {
	return r.db.Save(video).Error
}

func (r *videoRepository) Delete(id uuid.UUID) error {
	return r.db.Delete(&models.Video{}, "id = ?", id).Error
}

func (r *videoRepository) IncrementViews(id uuid.UUID) error {
	return r.db.Model(&models.Video{}).
		Where("id = ?", id).
		UpdateColumn("views", gorm.Expr("views + ?", 1)).Error
}

func (r *videoRepository) ExistsByID(id uuid.UUID) (bool, error) {
	var count int64
	err := r.db.Model(&models.Video{}).
		Where("id = ?", id).
		Count(&count).Error
	return count > 0, err
}

func (r *videoRepository) IsOwner(videoID, userID uuid.UUID) (bool, error) {
	var video models.Video
	err := r.db.Select("id").Where("id = ? AND user_id = ?", videoID, userID).First(&video).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
