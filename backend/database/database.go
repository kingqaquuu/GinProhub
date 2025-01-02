package database

import (
	"gin-prohub/models"
	"log"
	"time"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)
func InitDB(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NowFunc: func() time.Time {
            return time.Now().Local()
		},
	})
	if err!=nil{
		log.Fatal("数据库连接失败: %v", err)
	}
	db.AutoMigrate(&models.User{}, &models.Post{}, &models.Comment{})
	return db,nil
}