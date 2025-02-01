package main

import (
	"api-integration/config"
	"api-integration/models"
	"api-integration/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		panic("Error loading env")
	}

	db := config.ConnectDatabase()

	db.AutoMigrate(&models.User{}, &models.Post{})

	router := gin.Default()

	routes.SetupRoutes(router, db)

	router.Run(":8080")
}