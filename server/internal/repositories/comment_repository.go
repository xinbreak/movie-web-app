package repositories

import (
	"errors"

	"github.com/xinbreak/movie-web-app/internal/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CommentRepository interface {
	Create(comment *models.Comment) error
	Update(comment *models.Comment) error
	Delete(id uuid.UUID) error
	FindByID(id uuid.UUID) (*models.Comment, error)
	FindByVideoID(videoID uuid.UUID, parentID *uuid.UUID, page, pageSize int) ([]models.Comment, int64, error)
	FindReplies(commentID uuid.UUID, page, pageSize int) ([]models.Comment, int64, error)
	GetRepliesCount(commentID uuid.UUID) (int64, error)
	ExistsByID(id uuid.UUID) (bool, error)
	IsAuthor(commentID, userID uuid.UUID) (bool, error)
}

type commentRepository struct {
	db *gorm.DB
}

func NewCommentRepository(db *gorm.DB) CommentRepository {
	return &commentRepository{db: db}
}

func (r *commentRepository) Create(comment *models.Comment) error {
	return r.db.Create(comment).Error
}

func (r *commentRepository) Update(comment *models.Comment) error {
	return r.db.Save(comment).Error
}

func (r *commentRepository) Delete(id uuid.UUID) error {
	return r.db.Delete(&models.Comment{}, id).Error
}

func (r *commentRepository) FindByID(id uuid.UUID) (*models.Comment, error) {
	var comment models.Comment
	err := r.db.
		Preload("User").
		Preload("Video").
		Preload("Replies", func(db *gorm.DB) *gorm.DB {
			return db.Limit(3).Order("created_at DESC")
		}).
		Preload("Replies.User").
		First(&comment, id).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &comment, nil
}

func (r *commentRepository) FindByVideoID(videoID uuid.UUID, parentID *uuid.UUID, page, pageSize int) ([]models.Comment, int64, error) {
	var comments []models.Comment
	var total int64

	query := r.db.Model(&models.Comment{}).Where("video_id = ?", videoID)

	if parentID == nil {
		query = query.Where("parent_id IS NULL")
	} else {
		query = query.Where("parent_id = ?", parentID)
	}

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * pageSize
	err := query.
		Preload("User").
		Preload("Video").
		Order("created_at DESC").
		Offset(offset).
		Limit(pageSize).
		Find(&comments).Error

	if err != nil {
		return nil, 0, err
	}

	return comments, total, nil
}

func (r *commentRepository) FindReplies(commentID uuid.UUID, page, pageSize int) ([]models.Comment, int64, error) {
	var replies []models.Comment
	var total int64

	query := r.db.Model(&models.Comment{}).Where("parent_id = ?", commentID)

	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	offset := (page - 1) * pageSize
	err := query.
		Preload("User").
		Preload("Video").
		Order("created_at ASC").
		Offset(offset).
		Limit(pageSize).
		Find(&replies).Error

	if err != nil {
		return nil, 0, err
	}

	return replies, total, nil
}

func (r *commentRepository) GetRepliesCount(commentID uuid.UUID) (int64, error) {
	var count int64
	err := r.db.Model(&models.Comment{}).
		Where("parent_id = ?", commentID).
		Count(&count).Error
	return count, err
}

func (r *commentRepository) ExistsByID(id uuid.UUID) (bool, error) {
	var count int64
	err := r.db.Model(&models.Comment{}).
		Where("id = ?", id).
		Count(&count).Error
	return count > 0, err
}

func (r *commentRepository) IsAuthor(commentID, userID uuid.UUID) (bool, error) {
	var comment models.Comment
	err := r.db.Select("id").Where("id = ? AND user_id = ?", commentID, userID).First(&comment).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
