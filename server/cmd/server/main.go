package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/xinbreak/movie-web-app/database"
	_ "github.com/xinbreak/movie-web-app/docs"
	"github.com/xinbreak/movie-web-app/internal/controllers"
	"github.com/xinbreak/movie-web-app/internal/models"
	"github.com/xinbreak/movie-web-app/internal/repositories"
	"github.com/xinbreak/movie-web-app/internal/services"
)

// @title movie-web-app API
// @version 1.0
// @description API Server for Movie Application
// @host localhost:8080
// @BasePath /api
func main() {
	database.InitDB()
	database.DB.AutoMigrate(&models.User{}, &models.Video{}, &models.Comment{})

	userRepo := repositories.NewUserRepository(database.DB)
	userSvc := services.NewUserService(userRepo)
	userCtrl := controllers.NewUserController(userSvc)

	videoRepo := repositories.NewVideoRepository(database.DB)
	videoSvc := services.NewVideoService(videoRepo)
	videoCtrl := controllers.NewVideoController(videoSvc)

	commentRepo := repositories.NewCommentRepository(database.DB)
	commentSvc := services.NewCommentService(commentRepo, database.DB)
	commentCtrl := controllers.NewCommentController(commentSvc)

	r := gin.Default()
	r.Use(cors.Default())

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := r.Group("/api")
	{
		auth := api.Group("/auth")
		{
			auth.POST("/register", userCtrl.CreateUser)
			auth.POST("/login", userCtrl.Login)
		}

		users := api.Group("/users")
		{
			users.GET("", userCtrl.GetUsers)
			users.GET("/:id", userCtrl.GetUser)
			users.PUT("/:id", userCtrl.UpdateUser)
			users.DELETE("/:id", userCtrl.DeleteUser)
			users.GET("/:id/videos", videoCtrl.GetByUser)
		}

		videos := api.Group("/videos")
		{
			videos.GET("", videoCtrl.GetAll)
			videos.GET("/:id", videoCtrl.GetByID)
			videos.POST("", videoCtrl.Create)
			videos.PUT("/:id", videoCtrl.Update)
			videos.DELETE("/:id", videoCtrl.Delete)
			videos.GET("/:id/comments", commentCtrl.GetVideoComments)
		}

		comments := api.Group("/comments")
		{
			comments.POST("", commentCtrl.Create)
			comments.GET("/:id", commentCtrl.GetByID)
			comments.PUT("/:id", commentCtrl.Update)
			comments.DELETE("/:id", commentCtrl.Delete)
			comments.GET("/:id/replies", commentCtrl.GetCommentReplies)
		}
	}

	r.Run(":8080")
}
