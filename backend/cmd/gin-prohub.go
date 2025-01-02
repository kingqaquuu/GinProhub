package main

import (
	"gin-prohub/config"
	"gin-prohub/database"
	"gin-prohub/routes"
	"log"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)


func main() {
	var db *gorm.DB
	cfg,err:=config.LoadConfig()
	if err!=nil{
		log.Fatalf("配置加载失败: %v", err)
	}
	db, err = database.InitDB(cfg.DSN())
    if err != nil {
        log.Fatalf("数据库连接失败: %v", err)
    }

	router := gin.Default()
	routes.LoadRoutes(router,db)
	router.Static("/static", filepath.Join("..", "..", "frontend", "static"))
	router.LoadHTMLGlob(filepath.Join("..", "..", "frontend", "assets", "*"))
	router.GET("/", func(c *gin.Context) {
        c.HTML(http.StatusOK, "index.html", nil)
    })
	router.Run()
	
	//r.Run() // 监听并在 0.0.0.0:8080 上启动服务
}
