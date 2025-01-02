package routes

import (
	"gin-prohub/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)
func RegisterPostCmdRoutes(router *gin.Engine,db *gorm.DB){
	post:=router.Group("api/v1/posts")
	post.Use(middleware.JWTAuthMiddleware())
	{
		
	}
	
}