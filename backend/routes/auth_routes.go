package routes

import (
	"gin-prohub/models"
	"gin-prohub/services"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)
func RegisterAuthRoutes(router *gin.Engine,db *gorm.DB){
	users:=router.Group("api/v1/auth")
	{
		users.POST("/register",func(c *gin.Context){
			handleRegister(c,db)
		})
		users.POST("/login",)
	}
}

func handleRegister(c *gin.Context,db *gorm.DB){
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
	if exists := services.CheckUserExists(user.Username, user.Email,db); exists {
        c.JSON(http.StatusConflict, gin.H{"error": "Username or email already exists"})
        return
    }
	if err := user.SetPassword(user.Password); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to encrypt password"})
        return
    }
	user.LastLogin = time.Now()
	if err := db.Create(&user).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user"})
        return
    }
	c.JSON(http.StatusCreated, gin.H{
        "status":  "success",
        "message": "User registered successfully",
        "data":    user,
    })
}
