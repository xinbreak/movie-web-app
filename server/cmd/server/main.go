package main

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/xinbreak/movie-web-app/database"
	_ "github.com/xinbreak/movie-web-app/docs"
	"github.com/xinbreak/movie-web-app/internal/controllers"
	"github.com/xinbreak/movie-web-app/internal/models"
	"github.com/xinbreak/movie-web-app/internal/repository"
	"github.com/xinbreak/movie-web-app/internal/service"
)

// @title movie-web-app API
// @version 1.0
// @description API Server for Movie Application
// @host localhost:8080
// @BasePath /api
func main() {
	database.InitDB()
	database.DB.AutoMigrate(&models.User{})

	userRepo := repository.NewUserRepository(database.DB)
	userSvc := service.NewUserService(userRepo)
	userCtrl := controllers.NewUserController(userSvc)

	r := gin.Default()

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
		}
	}

	r.Run(":8080")
}
