package services

import (
	"errors"
	"gin-prohub/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CheckUserExists(username, email string, db *gorm.DB) bool {

	result := db.Where("username = ? OR email = ?", username, email).First(&models.User{})
	if result.Error == nil {
		return true
	}

	// 如果错误是记录未找到，则用户不存在
	if result.Error == gorm.ErrRecordNotFound {
		return false
	}
	return false
}

func RegisterUser(c *gin.Context, db *gorm.DB) error {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		return err
	}
	if exists := CheckUserExists(user.Username, user.Email, db); exists {
		return errors.New("username or email already exists")
	}
	if err := user.SetPassword(user.Password); err != nil {
		return err
	}

	user.LastLogin = time.Now()
	if err := db.Create(&user).Error; err != nil {
		return err
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  "success",
		"message": "User registered successfully",
		"data":    user,
	})
	return nil
}

func LoginUser(c *gin.Context, db *gorm.DB) error {
    // TODO: 实现登录逻辑
	
    return nil
}