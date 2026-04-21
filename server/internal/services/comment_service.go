package services

import (
	"errors"
	"fmt"

	"github.com/xinbreak/movie-web-app/internal/models"
	"github.com/xinbreak/movie-web-app/internal/repositories"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

var (
	ErrCommentNotFound      = errors.New("comment not found")
	ErrCommentEmptyContent  = errors.New("comment content cannot be empty")
	ErrInvalidParentComment = errors.New("invalid parent comment")
	ErrUnauthorized         = errors.New("you are not authorized to perform this action")
	ErrCommentTooLong       = errors.New("comment content exceeds maximum length")
)

type CommentService interface {
	CreateComment(comment *models.Comment) (*models.Comment, error)
	UpdateComment(id uuid.UUID, userID uuid.UUID, content string) (*models.Comment, error)
	DeleteComment(id uuid.UUID, userID uuid.UUID) error
	GetCommentByID(id uuid.UUID) (*models.Comment, error)
	GetVideoComments(videoID uuid.UUID, page, pageSize int) ([]models.Comment, int64, error)
	GetCommentReplies(commentID uuid.UUID, page, pageSize int) ([]models.Comment, int64, error)
	ValidateParentComment(videoID, parentID uuid.UUID) error
}

type commentService struct {
	commentRepo repositories.CommentRepository
	db          *gorm.DB
}

func NewCommentService(commentRepo repositories.CommentRepository, db *gorm.DB) CommentService {
	return &commentService{
		commentRepo: commentRepo,
		db:          db,
	}
}

func (s *commentService) CreateComment(comment *models.Comment) (*models.Comment, error) {
	if comment.Content == "" {
		return nil, ErrCommentEmptyContent
	}

	if len(comment.Content) > 5000 {
		return nil, ErrCommentTooLong
	}

	if comment.ParentID != nil {
		if err := s.ValidateParentComment(comment.VideoID, *comment.ParentID); err != nil {
			return nil, err
		}
	}

	if err := s.commentRepo.Create(comment); err != nil {
		return nil, fmt.Errorf("failed to create comment: %w", err)
	}

	_, err := s.commentRepo.FindByID(comment.ID)
	if err != nil {
		return nil, err
	}

	return comment, nil
}

func (s *commentService) UpdateComment(id, userID uuid.UUID, content string) (*models.Comment, error) {
	comment, err := s.commentRepo.FindByID(id)
	if err != nil {
		return nil, err
	}
	if comment == nil {
		return nil, ErrCommentNotFound
	}

	isAuthor, err := s.commentRepo.IsAuthor(id, userID)
	if err != nil {
		return nil, err
	}
	if !isAuthor {
		return nil, ErrUnauthorized
	}

	if content == "" {
		return nil, ErrCommentEmptyContent
	}
	if len(content) > 5000 {
		return nil, ErrCommentTooLong
	}

	comment.Content = content
	if err := s.commentRepo.Update(comment); err != nil {
		return nil, fmt.Errorf("failed to update comment: %w", err)
	}

	return comment, nil
}

func (s *commentService) DeleteComment(id, userID uuid.UUID) error {
	exists, err := s.commentRepo.ExistsByID(id)
	if err != nil {
		return err
	}
	if !exists {
		return ErrCommentNotFound
	}

	isAuthor, err := s.commentRepo.IsAuthor(id, userID)
	if err != nil {
		return err
	}
	if !isAuthor {
		return ErrUnauthorized
	}

	if err := s.commentRepo.Delete(id); err != nil {
		return fmt.Errorf("failed to delete comment: %w", err)
	}

	return nil
}

func (s *commentService) GetCommentByID(id uuid.UUID) (*models.Comment, error) {
	comment, err := s.commentRepo.FindByID(id)
	if err != nil {
		return nil, err
	}
	if comment == nil {
		return nil, ErrCommentNotFound
	}
	return comment, nil
}

func (s *commentService) GetVideoComments(videoID uuid.UUID, page, pageSize int) ([]models.Comment, int64, error) {

	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 20
	}

	comments, total, err := s.commentRepo.FindByVideoID(videoID, nil, page, pageSize)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to get video comments: %w", err)
	}

	for i := range comments {
		count, err := s.commentRepo.GetRepliesCount(comments[i].ID)
		if err != nil {
			continue
		}
		_ = count
	}

	return comments, total, nil
}

func (s *commentService) GetCommentReplies(commentID uuid.UUID, page, pageSize int) ([]models.Comment, int64, error) {
	exists, err := s.commentRepo.ExistsByID(commentID)
	if err != nil {
		return nil, 0, err
	}
	if !exists {
		return nil, 0, ErrCommentNotFound
	}

	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 50 {
		pageSize = 20
	}

	replies, total, err := s.commentRepo.FindReplies(commentID, page, pageSize)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to get comment replies: %w", err)
	}

	return replies, total, nil
}

func (s *commentService) ValidateParentComment(videoID, parentID uuid.UUID) error {
	parent, err := s.commentRepo.FindByID(parentID)
	if err != nil {
		return err
	}
	if parent == nil {
		return ErrInvalidParentComment
	}

	if parent.VideoID != videoID {
		return ErrInvalidParentComment
	}

	if parent.ParentID != nil {
		return errors.New("nested replies are not allowed (max depth: 1)")
	}

	return nil
}
