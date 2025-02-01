package controllers

import(
	"api-integration/models"

	"net/http"

	"github.com/gin-gonic/gin"

	"gorm.io/gorm"
)

type UserController struct {
	DB *gorm.DB
}

func NewUserController(db *gorm.DB) *UserController {
	return &UserController{DB: db}
}

func (c *UserController) GetUsers(ctx *gin.Context) {
	var users []models.User

	c.DB.Find(&users)

	ctx.JSON(http.StatusOK, gin.H{
		"data": users,
	})
}

func (c *UserController) GetUsersPost(ctx *gin.Context) {
	var users []models.User

	c.DB.Preload("Posts").Find(&users)

	ctx.JSON(http.StatusOK, gin.H{
		"data": users,
	})
}

func (c *UserController) GetUser(ctx *gin.Context) {
	id := ctx.Param("id")
	var user models.User

	result := c.DB.First(&user, id)

	if result.Error != nil {
		ctx.JSON(http.StatusNotFound, gin.H {
			"error": "User tidak ditemukan",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H {
		"data": user,
	})
}