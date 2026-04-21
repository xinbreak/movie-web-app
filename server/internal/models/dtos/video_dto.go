package dto

type CreateVideoRequest struct {
	Title       string `json:"title" binding:"required,min=1,max=255"`
	Description string `json:"description" binding:"omitempty,max=5000"`
	FilePath    string `json:"file_path" binding:"required"`
	Thumbnail   string `json:"thumbnail" binding:"omitempty,url"`
	Duration    int    `json:"duration" binding:"omitempty,min=0"`
}

type UpdateVideoRequest struct {
	Title       string `json:"title" binding:"required,min=1,max=255"`
	Description string `json:"description" binding:"omitempty,max=5000"`
}

type VideoResponseDTO struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Thumbnail   string `json:"thumbnail_url"`
	Duration    int    `json:"duration"`
	ViewsCount  uint32 `json:"views_count"`
	UserID      string `json:"user_id"`
	Status      string `json:"status"`
}
