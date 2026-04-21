package controllers

import (
	"net/http"
	"strconv"

	"github.com/xinbreak/movie-web-app/internal/models"
	dto "github.com/xinbreak/movie-web-app/internal/models/dtos"
	"github.com/xinbreak/movie-web-app/internal/services"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CommentController struct {
	commentService services.CommentService
}

func NewCommentController(commentService services.CommentService) *CommentController {
	return &CommentController{
		commentService: commentService,
	}
}

// Create создает новый комментарий
// @Summary Create a new comment
// @Description Create a new comment for a video
// @Tags comments
// @Accept json
// @Produce json
// @Param request body dto.CreateCommentRequest true "Comment creation request"
// @Success 201 {object} models.Comment
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /comments [post]
func (ctrl *CommentController) Create(c *gin.Context) {
	var req dto.CreateCommentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "invalid request",
			"details": err.Error(),
		})
		return
	}

	// Получаем userID из контекста (после аутентификации)
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "unauthorized",
		})
		return
	}

	// Преобразуем userID в uuid.UUID
	userUUID, ok := userID.(uuid.UUID)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "invalid user id format",
		})
		return
	}

	comment := &models.Comment{
		Content: req.Content,
		UserID:  userUUID,
		//VideoID:  req.PostID,
		ParentID: req.ParentID,
	}

	created, err := ctrl.commentService.CreateComment(comment)
	if err != nil {
		switch err {
		case services.ErrCommentEmptyContent:
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "comment content cannot be empty",
			})
		case services.ErrCommentTooLong:
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "comment content exceeds maximum length of 5000 characters",
			})
		case services.ErrInvalidParentComment:
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "invalid parent comment",
			})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "failed to create comment",
			})
		}
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "comment created successfully",
		"comment": created,
	})
}

// Update обновляет комментарий
// @Summary Update a comment
// @Description Update an existing comment by ID
// @Tags comments
// @Accept json
// @Produce json
// @Param id path string true "Comment ID"
// @Param request body dto.UpdateCommentRequest true "Comment update request"
// @Success 200 {object} models.Comment
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 403 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /comments/{id} [put]
func (ctrl *CommentController) Update(c *gin.Context) {
	idStr := c.Param("id")
	commentID, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid comment id",
		})
		return
	}

	var req dto.UpdateCommentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "invalid request",
			"details": err.Error(),
		})
		return
	}

	// Получаем userID из контекста
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "unauthorized",
		})
		return
	}

	userUUID, ok := userID.(uuid.UUID)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "invalid user id format",
		})
		return
	}

	updated, err := ctrl.commentService.UpdateComment(commentID, userUUID, req.Content)
	if err != nil {
		switch err {
		case services.ErrCommentNotFound:
			c.JSON(http.StatusNotFound, gin.H{
				"error": "comment not found",
			})
		case services.ErrUnauthorized:
			c.JSON(http.StatusForbidden, gin.H{
				"error": "you don't have permission to update this comment",
			})
		case services.ErrCommentEmptyContent:
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "comment content cannot be empty",
			})
		case services.ErrCommentTooLong:
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "comment content exceeds maximum length",
			})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "failed to update comment",
			})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "comment updated successfully",
		"comment": updated,
	})
}

// Delete удаляет комментарий
// @Summary Delete a comment
// @Description Delete a comment by ID (only author can delete)
// @Tags comments
// @Produce json
// @Param id path string true "Comment ID"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 403 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /comments/{id} [delete]
func (ctrl *CommentController) Delete(c *gin.Context) {
	idStr := c.Param("id")
	commentID, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid comment id",
		})
		return
	}

	// Получаем userID из контекста
	userID, exists := c.Get("user_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "unauthorized",
		})
		return
	}

	userUUID, ok := userID.(uuid.UUID)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "invalid user id format",
		})
		return
	}

	err = ctrl.commentService.DeleteComment(commentID, userUUID)
	if err != nil {
		switch err {
		case services.ErrCommentNotFound:
			c.JSON(http.StatusNotFound, gin.H{
				"error": "comment not found",
			})
		case services.ErrUnauthorized:
			c.JSON(http.StatusForbidden, gin.H{
				"error": "you don't have permission to delete this comment",
			})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "failed to delete comment",
			})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "comment deleted successfully",
	})
}

// GetByID получает комментарий по ID
// @Summary Get a comment by ID
// @Description Get a single comment by its ID with replies
// @Tags comments
// @Produce json
// @Param id path string true "Comment ID"
// @Success 200 {object} models.Comment
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /comments/{id} [get]
func (ctrl *CommentController) GetByID(c *gin.Context) {
	idStr := c.Param("id")
	commentID, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid comment id",
		})
		return
	}

	comment, err := ctrl.commentService.GetCommentByID(commentID)
	if err != nil {
		switch err {
		case services.ErrCommentNotFound:
			c.JSON(http.StatusNotFound, gin.H{
				"error": "comment not found",
			})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "failed to get comment",
			})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"comment": comment,
	})
}

// GetVideoComments получает комментарии для видео
// @Summary Get video comments
// @Description Get paginated comments for a specific video
// @Tags comments
// @Produce json
// @Param video_id path string true "Video ID"
// @Param page query int false "Page number" default(1)
// @Param page_size query int false "Items per page" default(20)
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /videos/{video_id}/comments [get]
func (ctrl *CommentController) GetVideoComments(c *gin.Context) {
	videoIDStr := c.Param("video_id")
	videoID, err := uuid.Parse(videoIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid video id",
		})
		return
	}

	// Парсинг параметров пагинации
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil || page < 1 {
		page = 1
	}

	pageSize, err := strconv.Atoi(c.DefaultQuery("page_size", "20"))
	if err != nil || pageSize < 1 {
		pageSize = 20
	}
	if pageSize > 100 {
		pageSize = 100
	}

	comments, total, err := ctrl.commentService.GetVideoComments(videoID, page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to get comments",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"comments":    comments,
		"total":       total,
		"page":        page,
		"page_size":   pageSize,
		"total_pages": (total + int64(pageSize) - 1) / int64(pageSize),
	})
}

// GetCommentReplies получает ответы на комментарий
// @Summary Get comment replies
// @Description Get paginated replies for a specific comment
// @Tags comments
// @Produce json
// @Param id path string true "Comment ID"
// @Param page query int false "Page number" default(1)
// @Param page_size query int false "Items per page" default(20)
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /comments/{id}/replies [get]
func (ctrl *CommentController) GetCommentReplies(c *gin.Context) {
	idStr := c.Param("id")
	commentID, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid comment id",
		})
		return
	}

	// Парсинг параметров пагинации
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil || page < 1 {
		page = 1
	}

	pageSize, err := strconv.Atoi(c.DefaultQuery("page_size", "20"))
	if err != nil || pageSize < 1 {
		pageSize = 20
	}
	if pageSize > 50 {
		pageSize = 50
	}

	replies, total, err := ctrl.commentService.GetCommentReplies(commentID, page, pageSize)
	if err != nil {
		switch err {
		case services.ErrCommentNotFound:
			c.JSON(http.StatusNotFound, gin.H{
				"error": "comment not found",
			})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "failed to get replies",
			})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"replies":     replies,
		"total":       total,
		"page":        page,
		"page_size":   pageSize,
		"total_pages": (total + int64(pageSize) - 1) / int64(pageSize),
	})
}
