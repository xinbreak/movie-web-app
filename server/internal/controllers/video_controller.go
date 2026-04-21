package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/xinbreak/movie-web-app/internal/models"
	dto "github.com/xinbreak/movie-web-app/internal/models/dtos"
	"github.com/xinbreak/movie-web-app/internal/services"
)

type VideoController struct {
	videoService services.VideoService
}

func NewVideoController(videoService services.VideoService) *VideoController {
	return &VideoController{
		videoService: videoService,
	}
}

// Create создает новое видео
// @Summary Create a new video
// @Description Upload a new video with title and description
// @Tags videos
// @Accept json
// @Produce json
// @Param request body dto.CreateVideoRequest true "Video creation request"
// @Success 201 {object} models.Video
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /videos [post]
func (ctrl *VideoController) Create(c *gin.Context) {
	var req dto.CreateVideoRequest
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

	userUUID, ok := userID.(uuid.UUID)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "invalid user id format",
		})
		return
	}

	video := &models.Video{
		Title:       req.Title,
		Description: req.Description,
		FilePath:    req.FilePath,
		Thumbnail:   req.Thumbnail,
		Duration:    req.Duration,
		UserID:      userUUID,
		Status:      "uploaded",
	}

	created, err := ctrl.videoService.CreateVideo(video)
	if err != nil {
		switch err {
		case services.ErrVideoEmptyTitle:
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "video title cannot be empty",
			})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "failed to create video",
			})
		}
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "video created successfully",
		"video":   created,
	})
}

// GetByID получает видео по ID
// @Summary Get video by ID
// @Description Get a single video by its ID
// @Tags videos
// @Produce json
// @Param id path string true "Video ID"
// @Success 200 {object} models.Video
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /videos/{id} [get]
func (ctrl *VideoController) GetByID(c *gin.Context) {
	idStr := c.Param("id")
	videoID, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid video id",
		})
		return
	}

	video, err := ctrl.videoService.GetVideoByID(videoID)
	if err != nil {
		switch err {
		case services.ErrVideoNotFound:
			c.JSON(http.StatusNotFound, gin.H{
				"error": "video not found",
			})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "failed to get video",
			})
		}
		return
	}

	// Increment views
	go ctrl.videoService.IncrementViews(videoID)

	c.JSON(http.StatusOK, gin.H{
		"video": video,
	})
}

// GetAll получает все видео с пагинацией
// @Summary Get all videos
// @Description Get paginated list of all videos
// @Tags videos
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param page_size query int false "Items per page" default(20)
// @Success 200 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /videos [get]
func (ctrl *VideoController) GetAll(c *gin.Context) {
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

	videos, total, err := ctrl.videoService.GetAllVideos(page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to get videos",
		})
		return
	}

	totalPages := int(total) / pageSize
	if int(total)%pageSize > 0 {
		totalPages++
	}

	c.JSON(http.StatusOK, gin.H{
		"videos":      videos,
		"total":       total,
		"page":        page,
		"page_size":   pageSize,
		"total_pages": totalPages,
	})
}

// GetByUser получает видео конкретного пользователя
// @Summary Get videos by user
// @Description Get paginated list of videos for a specific user
// @Tags videos
// @Produce json
// @Param user_id path string true "User ID"
// @Param page query int false "Page number" default(1)
// @Param page_size query int false "Items per page" default(20)
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /users/{user_id}/videos [get]
func (ctrl *VideoController) GetByUser(c *gin.Context) {
	userIDStr := c.Param("user_id")
	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid user id",
		})
		return
	}

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

	videos, total, err := ctrl.videoService.GetVideosByUser(userID, page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "failed to get user videos",
		})
		return
	}

	totalPages := int(total) / pageSize
	if int(total)%pageSize > 0 {
		totalPages++
	}

	c.JSON(http.StatusOK, gin.H{
		"videos":      videos,
		"total":       total,
		"page":        page,
		"page_size":   pageSize,
		"total_pages": totalPages,
	})
}

// Update обновляет видео
// @Summary Update a video
// @Description Update an existing video (only owner can update)
// @Tags videos
// @Accept json
// @Produce json
// @Param id path string true "Video ID"
// @Param request body dto.UpdateVideoRequest true "Video update request"
// @Success 200 {object} models.Video
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 403 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /videos/{id} [put]
func (ctrl *VideoController) Update(c *gin.Context) {
	idStr := c.Param("id")
	videoID, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid video id",
		})
		return
	}

	var req dto.UpdateVideoRequest
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

	updated, err := ctrl.videoService.UpdateVideo(videoID, userUUID, req.Title, req.Description)
	if err != nil {
		switch err {
		case services.ErrVideoNotFound:
			c.JSON(http.StatusNotFound, gin.H{
				"error": "video not found",
			})
		case services.ErrUnauthorizedVideo:
			c.JSON(http.StatusForbidden, gin.H{
				"error": "you don't have permission to update this video",
			})
		case services.ErrVideoEmptyTitle:
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "video title cannot be empty",
			})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "failed to update video",
			})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "video updated successfully",
		"video":   updated,
	})
}

// Delete удаляет видео
// @Summary Delete a video
// @Description Delete a video by ID (only owner can delete)
// @Tags videos
// @Produce json
// @Param id path string true "Video ID"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Failure 403 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /videos/{id} [delete]
func (ctrl *VideoController) Delete(c *gin.Context) {
	idStr := c.Param("id")
	videoID, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid video id",
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

	err = ctrl.videoService.DeleteVideo(videoID, userUUID)
	if err != nil {
		switch err {
		case services.ErrVideoNotFound:
			c.JSON(http.StatusNotFound, gin.H{
				"error": "video not found",
			})
		case services.ErrUnauthorizedVideo:
			c.JSON(http.StatusForbidden, gin.H{
				"error": "you don't have permission to delete this video",
			})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "failed to delete video",
			})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "video deleted successfully",
	})
}
