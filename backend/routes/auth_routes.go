package routes

import (
	"gin-prohub/services"
	"gin-prohub/utils/response"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func RegisterAuthRoutes(router *gin.Engine, db *gorm.DB) {
	auth := router.Group("api/v1/auth")
	{
		auth.POST("/register", handleRegister(db)) //注册
		auth.POST("/login", handleLogin(db))       //登录
		auth.POST("/logout", handleLogout())       //登出
	}
}

func handleRegister(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		if err := services.RegisterUser(c, db); err != nil {
			response.Error(c, http.StatusBadRequest, err.Error())
			return
		}
		response.Success(c, gin.H{
			"message": "User registered successfully",
		})
	}
}
func handleLogin(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		data,err:=services.LoginUser(c,db)
		if err!=nil{
			response.Error(c,http.StatusBadRequest,err.Error())
		}
		response.Success(c, data)
	}
}

func handleLogout() gin.HandlerFunc {
	return func(c *gin.Context) {
		response.Success(c, gin.H{
			"message": "Logout successful",
		})
	}
}
