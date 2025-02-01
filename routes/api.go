package routes

import (
	"api-integration/controllers"

	"github.com/gin-gonic/gin"

	"gorm.io/gorm"
)

func SetupRoutes(router *gin.Engine, db *gorm.DB) {
	userController := controllers.NewUserController(db)

	api := router.Group("/api")
	{
		api.GET("/users", userController.GetUsers)
		api.GET("/users/:id", userController.GetUser)
		api.GET("/posts", userController.GetUsersPost)
	}
	
}