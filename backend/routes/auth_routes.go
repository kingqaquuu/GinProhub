package routes

import (
	"gin-prohub/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)
func RegisterAuthRoutes(router *gin.Engine,db *gorm.DB){
	auth:=router.Group("api/v1/auth")
	{
		auth.POST("/register", handleRegister(db))//注册
        auth.POST("/login", handleLogin(db))//登录
        auth.POST("/logout", handleLogout())//登出
	}
}

func handleRegister(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
        if err := services.RegisterUser(c, db); err != nil {
            response.Error(c, http.StatusBadRequest, err.Error())
            return
        }
    }
}
func handleLogin(db *gorm.DB) gin.HandlerFunc {
    return func(c *gin.Context) {
        if err := services.LoginUser(c, db); err != nil {
            c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
            return
        }
    }
}

func handleLogout() gin.HandlerFunc {
    return func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{"message": "Successfully logged out"})
    }
}